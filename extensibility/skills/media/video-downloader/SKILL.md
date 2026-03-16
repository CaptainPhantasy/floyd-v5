---
name: video-downloader
description: Downloads videos from YouTube and other platforms for offline viewing, editing, or archival. Handles various formats and quality options using yt-dlp. Use when user wants to download videos, save content offline, or extract audio from videos.
---

# Video Downloader

Wraps yt-dlp workflows to download videos and audio from YouTube and other platforms.

## Capabilities

- Download single videos, playlists, or batches
- Select quality (480p—4K)
- Download audio-only (e.g., MP3)
- Preserve basic metadata and thumbnails

## Installation

```bash
pip install yt-dlp
```

## Common Usage Patterns

### Download Best Quality Video

```bash
yt-dlp "https://www.youtube.com/watch?v=VIDEO_ID"
```

### Download Specific Quality

```bash
# 1080p
yt-dlp -f "bestvideo[height<=1080]+bestaudio" "URL"

# 720p
yt-dlp -f "bestvideo[height<=720]+bestaudio" "URL"

# 480p
yt-dlp -f "bestvideo[height<=480]+bestaudio" "URL"
```

### Audio Only (MP3)

```bash
yt-dlp -x --audio-format mp3 "URL"
```

### Download Playlist

```bash
yt-dlp -o "%(playlist_index)s-%(title)s.%(ext)s" "PLAYLIST_URL"
```

### Preserve Metadata and Thumbnail

```bash
yt-dlp --embed-thumbnail --embed-metadata "URL"
```

### Download Multiple Videos from File

```bash
# Create urls.txt with one URL per line
yt-dlp -a urls.txt
```

## Quality Options

```bash
# List available formats
yt-dlp -F "URL"

# Download specific format
yt-dlp -f FORMAT_CODE "URL"
```

## Output Templates

```bash
# Custom filename
yt-dlp -o "%(title)s.%(ext)s" "URL"

# Include upload date
yt-dlp -o "%(upload_date)s-%(title)s.%(ext)s" "URL"

# Organize by uploader
yt-dlp -o "%(uploader)s/%(title)s.%(ext)s" "URL"
```

## Supported Platforms

yt-dlp supports 1000+ sites including:
- YouTube
- Vimeo
- Twitter/X
- TikTok
- Reddit
- Facebook
- Instagram
- And many more

Check support: `yt-dlp --list-extractors`

## Copyright and Fair Use

**Important:** Downloading videos may violate copyright laws or platform terms of service.

Users are responsible for:
- Complying with copyright laws in their jurisdiction
- Respecting platform terms of service
- Using downloads only for permitted purposes (personal backup, fair use, etc.)

This skill is provided for lawful use only. Always verify you have the right to download content before doing so.

## Common Issues

**Download fails:**
```bash
# Update yt-dlp
pip install --upgrade yt-dlp
```

**Geo-restricted content:**
```bash
yt-dlp --geo-bypass "URL"
```

**Age-restricted content:**
```bash
yt-dlp --cookies cookies.txt "URL"
```
