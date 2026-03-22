/**
 * HTTP Server for FLOYD SUPERCACHE
 *
 * Simple Express server providing REST API access to SUPERCACHE
 * alongside the MCP stdio interface.
 */

import express from 'express';
import cors from 'cors';
import { readFileSync, existsSync, writeFileSync, mkdirSync } from 'fs';
import { join, dirname } from 'path';
import { homedir } from 'os';
import { fileURLToPath } from 'url';
import { lockSync, unlockSync } from 'proper-lockfile';

const __dirname = dirname(fileURLToPath(import.meta.url));

// Cache storage paths (same as MCP server)
const CACHE_DIR = join(homedir(), '.floyd', 'supercache');
const INDEX_FILE = join(CACHE_DIR, 'index.json');

interface CacheEntry {
  key: string;
  namespace: string;
  value: any;
  tier: 'project' | 'reasoning' | 'vault';
  createdAt: string;
  expiresAt?: string;
  accessCount: number;
  lastAccessed: string;
  tags: string[];
  metadata?: Record<string, any>;
}

interface CacheIndex {
  entries: Record<string, CacheEntry>;
}

function loadIndex(): CacheIndex {
  if (existsSync(INDEX_FILE)) {
    try {
      lockSync(INDEX_FILE);
      try {
        const data = readFileSync(INDEX_FILE, 'utf8');
        return JSON.parse(data);
      } finally {
        unlockSync(INDEX_FILE);
      }
    } catch {
      return { entries: {} };
    }
  }
  return { entries: {} };
}

function saveIndex(index: CacheIndex): void {
  lockSync(INDEX_FILE);
  try {
    writeFileSync(INDEX_FILE, JSON.stringify(index, null, 2), 'utf8');
  } finally {
    unlockSync(INDEX_FILE);
  }
}

function sanitizeKey(key: string): string {
  return key.replace(/[^a-zA-Z0-9_-]/g, '_');
}

function getEntryKey(key: string, namespace: string = 'global'): string {
  return `${sanitizeKey(namespace)}:${sanitizeKey(key)}`;
}

// Create Express app
const app = express();
app.use(cors());
app.use(express.json());

// Request logging middleware to track what's calling the server
app.use((req, res, next) => {
  const timestamp = new Date().toISOString();
  const userAgent = req.get('User-Agent') || 'Unknown';
  const referer = req.get('Referer') || 'Direct';
  
  console.error(`[${timestamp}] ${req.method} ${req.url}`);
  console.error(`  User-Agent: ${userAgent}`);
  console.error(`  Referer: ${referer}`);
  console.error(`  IP: ${req.ip || req.connection.remoteAddress}`);
  console.error(`  Headers:`, JSON.stringify(req.headers, null, 2));
  console.error(`  Body:`, JSON.stringify(req.body, null, 2));
  console.error('---');
  
  next();
});

// GET /supercache/get?key=<key>
app.get('/supercache/get', (req, res) => {
  const { key, namespace = 'global' } = req.query;
  
  if (!key || typeof key !== 'string') {
    return res.status(400).json({ error: 'key parameter is required' });
  }

  try {
    const index = loadIndex();
    
    // Handle both formats: "system:project_registry" or separate namespace/key
    let entryKey: string;
    const keyStr = key as string;
    if (keyStr.includes(':')) {
      // If key contains namespace, use it directly
      entryKey = sanitizeKey(keyStr);
    } else {
      // Otherwise use separate namespace parameter
      entryKey = getEntryKey(keyStr, namespace as string);
    }
    
    const entry = index.entries[entryKey];

    if (!entry) {
      return res.status(404).json({ error: 'Key not found' });
    }

    // Update access stats
    entry.accessCount++;
    entry.lastAccessed = new Date().toISOString();
    saveIndex(index);

    res.json({
      key: entry.key,
      namespace: entry.namespace,
      value: entry.value,
      tier: entry.tier,
      createdAt: entry.createdAt,
      accessCount: entry.accessCount,
      lastAccessed: entry.lastAccessed,
      tags: entry.tags,
      metadata: entry.metadata
    });
  } catch (error) {
    console.error('Error getting cache entry:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// POST /supercache/set
app.post('/supercache/set', (req, res) => {
  const { key, value, namespace = 'global', tier = 'project', tags = [], metadata = {} } = req.body;

  if (!key || value === undefined) {
    return res.status(400).json({ error: 'key and value are required' });
  }

  try {
    const index = loadIndex();
    
    // Handle both formats: "system:project_registry" or separate namespace/key
    let entryKey: string;
    let actualNamespace = namespace;
    let actualKey = key;
    
    if (key.includes(':')) {
      // If key contains namespace, extract it
      const parts = key.split(':', 2);
      actualNamespace = parts[0];
      actualKey = parts[1];
      entryKey = sanitizeKey(`${actualNamespace}:${actualKey}`);
    } else {
      entryKey = getEntryKey(actualKey, actualNamespace);
    }
    
    const now = new Date().toISOString();

    const entry: CacheEntry = {
      key: actualKey,
      namespace: actualNamespace,
      value,
      tier,
      createdAt: now,
      accessCount: 0,
      lastAccessed: now,
      tags,
      metadata
    };

    index.entries[entryKey] = entry;
    saveIndex(index);

    res.json({ success: true, key: entryKey });
  } catch (error) {
    console.error('Error setting cache entry:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// GET /supercache/list
app.get('/supercache/list', (req, res) => {
  const { namespace = 'global', tier } = req.query;

  try {
    const index = loadIndex();
    let entries = Object.values(index.entries);

    // Filter by namespace
    if (namespace && namespace !== 'all') {
      entries = entries.filter(entry => entry.namespace === namespace);
    }

    // Filter by tier
    if (tier && tier !== 'all') {
      entries = entries.filter(entry => entry.tier === tier);
    }

    const summary = entries.map(entry => ({
      key: entry.key,
      namespace: entry.namespace,
      tier: entry.tier,
      createdAt: entry.createdAt,
      accessCount: entry.accessCount,
      lastAccessed: entry.lastAccessed,
      tags: entry.tags
    }));

    res.json({ entries: summary, count: summary.length });
  } catch (error) {
    console.error('Error listing cache entries:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// DELETE /supercache/delete
app.delete('/supercache/delete', (req, res) => {
  const { key, namespace = 'global' } = req.query;

  if (!key || typeof key !== 'string') {
    return res.status(400).json({ error: 'key parameter is required' });
  }

  try {
    const index = loadIndex();
    const entryKey = getEntryKey(key as string, namespace as string);
    
    if (!index.entries[entryKey]) {
      return res.status(404).json({ error: 'Key not found' });
    }

    delete index.entries[entryKey];
    saveIndex(index);

    res.json({ success: true, deleted: entryKey });
  } catch (error) {
    console.error('Error deleting cache entry:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// Health check
app.get('/health', (req, res) => {
  res.json({ status: 'ok', service: 'floyd-supercache-http' });
});

export async function startSupercacheHttpServer(port: number = 16999): Promise<void> {
  // Ensure cache directory exists
  if (!existsSync(CACHE_DIR)) {
    mkdirSync(CACHE_DIR, { recursive: true });
  }

  app.listen(port, '127.0.0.1', () => {
    console.error(`FLOYD SUPERCACHE HTTP Server running on http://127.0.0.1:${port}`);
  });
}

// Start server if this file is run directly
if (import.meta.url === `file://${process.argv[1]}`) {
  const port = parseInt(process.argv[2]) || 16999;
  startSupercacheHttpServer(port).catch(console.error);
}