name: "psi-project-expert"
description: "Expert knowledge of Precision Sewer Inspection project - business model, tech stack, database schema, PWA workflows, and development context."
trigger: "psi"
version: "1.0.0"
tags: ["project", "sewer-inspection", "pwa", "nextjs", "prisma"]
---

# PSI Project Expert

You are an expert on the Precision Sewer Inspection (PSI) project. Apply this knowledge to any task involving this codebase.

## PROJECT IDENTITY

**Business:** Precision Sewer Inspection - HD sewer scope inspections for Central Indiana  
**URL:** https://precisionsewerinspections.com  
**Location:** 6405 Justins Ridge Road, Nashville, IN 47448  
**Contact:** (317) 620-3858 | booking@precisionsewerinspections.com  
**Certification:** InterNACHI Certified  
**Core Policy:** NO UPSELLING - inspectors are not contractors

## BUSINESS MODEL

### Services
- Sewer scope inspection: $159 base (cleanout access)
- Same-day delivery: +$39
- Additional units (multi-family): $129 each
- Video Review service: FREE

### Access Method Fees
| Method | Fee |
|--------|-----|
| Standard Cleanout | Included |
| Roof Vent Access | +$50 |
| Toilet Pull & Reset | +$65 |
| Cleanout Cap Replacement | +$50 |
| Crawl Space Access | +$30 |
| Trip Fee (no access) | $79 |

### Volume Packages
- 10-Scope Bundle: $135/scope (~15% savings)
- 25-Scope Brokerage: Call for pricing
- Enterprise (400-600+/year): Call for pricing

### Service Area
Indianapolis, Carmel, Fishers, Noblesville, Westfield, Zionsville, Brownsburg, Avon, Plainfield, Greenwood, Franklin, Greenfield

### Delivery SLA
- Standard: 24 hours
- Same-day: +$39

### Active Promotion
Code: SAVE10 ($10 off first inspection)

## TECH STACK

| Layer | Technology |
|-------|------------|
| Framework | Next.js 14 (App Router) |
| Language | TypeScript |
| Styling | Tailwind CSS |
| Components | Radix UI + shadcn/ui |
| State | Zustand + React Query |
| Database | PostgreSQL + Prisma ORM |
| Cloud Storage | AWS S3 |
| Payments | Stripe (LIVE mode) |
| Scheduling | Google Calendar API |
| Maps | Mapbox GL |
| Notifications | Abacus (email) |

## PROJECT STRUCTURE

```
/Volumes/Storage/PSI/
├── FLOYD.md                          # Protocol file
├── HANDOFF.md                        # Session continuity
├── PROJECT_REFERENCE.md              # This knowledge base
├── Floyd PWA Plan.md                 # Original exploration
├── precision_sewer_backup/           # Database CSVs, schema, env
│   ├── database/                     # 15 CSV exports
│   ├── schema.prisma
│   └── env-backup.txt
├── precision_sewer_inspection (1)/
│   └── nextjs_space/                 # MAIN CODEBASE
│       ├── app/
│       │   ├── page.tsx              # Home page
│       │   ├── admin/                # Admin dashboard
│       │   ├── technician/           # Technician PWA
│       │   ├── download/             # Client portal
│       │   ├── status/               # Status tracking
│       │   ├── booking/              # Booking flow
│       │   └── api/                  # API routes
│       ├── components/
│       │   ├── inspection/           # 15 stage components
│       │   └── ui/                   # UI primitives
│       ├── lib/
│       │   ├── constants.ts          # Business constants
│       │   ├── db.ts                 # Prisma client
│       │   └── aws-config.ts         # S3 config
│       └── prisma/
│           └── schema.prisma         # Database schema
└── S830ASMKT_User_Manual-EN.pdf      # Camera equipment manual
```

## DATABASE SCHEMA

### Core Entities
- **User** - Technicians, admins, super_admins
- **Job** - Booking/request before assignment
- **Inspection** - Active work record with 12 stages
- **VideoAttachment** - Uploaded video with chapters
- **DeliveryToken** - 72-hour client access links
- **GeneratedReport** - AI-assisted reports
- **InspectionPhoto** - Photo evidence
- **LocationLog** - GPS tracking
- **VoiceRecording** - Field notes with transcription
- **ClientSignature** - Digital sign-off

### Key Enums

**InspectionStage (workflow order):**
```
ACCEPTED → EN_ROUTE → ARRIVED → PRE_INSPECTION → INSPECTING → 
POST_INSPECTION → VIDEO_ATTACH → CLIENT_SIGNOFF → SUBMITTED → 
UNDER_REVIEW → APPROVED → DELIVERED
```

**JobStatus:**
```
PENDING → ASSIGNED → ACCEPTED → EN_ROUTE → ON_SITE → IN_PROGRESS → COMPLETED → CANCELLED
```

**ConditionRating:** GOOD | FAIR | NEEDS_ATTENTION | CRITICAL  
**UrgencyLevel:** NONE | MONITOR | SOON | IMMEDIATE  
**AccessType:** CLEANOUT | ROOF_VENT | TOILET_PULL | UNKNOWN

### Key Relationships
```
Job 1:1 Inspection
Inspection 1:N LocationLogs (GPS)
Inspection 1:N Photos
Inspection 1:1 VideoAttachment → 1:N VideoChapters
Inspection 1:N VoiceRecordings
Inspection 1:1 ClientSignature
Inspection 1:1 GeneratedReport
Inspection 1:1 DeliveryToken
User 1:N Inspections (technician)
```

## PWA WORKFLOWS

### Technician PWA (`/technician`)

**9 Stage Flow:**
1. **ACCEPTED** - View job, tap "En Route"
2. **EN_ROUTE** - GPS tracking, navigation
3. **ARRIVED** - Photo verify, confirm access method
4. **PRE_INSPECTION** - Client interview (home age, issues, history)
5. **INSPECTING** - Camera scope, video recording, voice notes
6. **POST_INSPECTION** - Mark findings, select condition/urgency
7. **VIDEO_ATTACH** - Upload video, verify integrity
8. **CLIENT_SIGNOFF** - Capture digital signature
9. **SUBMITTED** - Submit for admin review

**Stage Components:**
- `en-route-stage.tsx` - GPS tracking
- `arrived-stage.tsx` - Arrival confirmation
- `pre-inspection-stage.tsx` - Client interview
- `inspecting-stage.tsx` - Scope execution
- `post-inspection-stage.tsx` - Findings entry
- `video-attach-stage.tsx` - Video upload
- `signature-stage.tsx` - Digital signature
- `submit-stage.tsx` - Final submission

**Supporting Components:**
- `camera-pairing.tsx` - Hardware integration
- `voice-listener.tsx` - Voice recording/transcription
- `video-chapters.tsx` - Chapter markers
- `highlight-reel.tsx` - Auto highlight generation
- `ai-summary.tsx` - AI report assistance
- `property-lookup.tsx` - Auto property data
- `field-validator.tsx` - Checklist validation

### Admin Dashboard (`/admin`)
- Filter by status: SUBMITTED, UNDER_REVIEW, APPROVED, REJECTED
- Video review with chapter navigation
- Report generation/editing
- Delivery link creation (72-hour expiration)
- Client notification triggering

### Client Portal (`/download/[inspectionId]`)
- HD video streaming
- Clickable issue markers (chapters)
- Video download
- PDF report download
- 72-hour access expiration countdown

### Status Tracking (`/status`)
- Input: Job Number + Email
- Output: Progress timeline, download link when ready

## HARD GATES

| Gate | Stage | Requirement |
|------|-------|-------------|
| 1.1 | Dispatch | Accept within 15 minutes |
| 2.1 | Arrival | GPS within 100ft of address |
| 2.2 | Arrival | On-time arrival logged |
| 3.1 | Site Verify | Access method + photo confirmed |
| 3.2 | Site Verify | Trip fee authorization if no access |
| 4.1 | Inspection | Video valid, minimum 30 seconds |
| 4.2 | Inspection | Overall condition selected |
| 5.1 | Upload | Server confirms file integrity |
| 5.2 | Upload | Field notes submitted |
| 6.1 | Sign-off | All checklist items complete |

## ISSUE TYPES

- Root Intrusion - Tree roots in pipe
- Cracks & Breaks - Structural damage
- Belly/Sag - Low spot collecting waste
- Blockages - Obstruction
- Scale Buildup - Mineral deposits
- Offset Joints - Misaligned sections
- Collapsed Pipe - Total failure
- Grease Buildup - FOG accumulation
- Debris/Foreign - Objects in line
- Tap Connection - Secondary line point
- Clean - No issues

## API ROUTES

### Technician
- `GET /api/technician/inspections` - List assigned
- `GET /api/technician/inspections/[id]` - Get details
- `POST /api/technician/inspections/[id]/stage` - Update stage
- `PATCH /api/technician/inspections/[id]` - Update data
- `GET /api/technician/jobs` - Available jobs
- `POST /api/technician/jobs/[id]/accept` - Accept job

### Admin
- `GET /api/admin/inspections` - List by status
- `GET /api/admin/inspections/[id]` - Review details
- `POST /api/admin/inspections/[id]/approve` - Approve
- `POST /api/admin/inspections/[id]/reject` - Return for corrections

### Client
- `GET /api/status/[jobNumber]?email=X` - Check status
- `GET /api/download/[inspectionId]` - Access portal

## INTEGRATIONS

### Stripe
- Mode: LIVE (pk_live_...)
- Products: Inspection services
- Promo codes: SAVE10

### Google Calendar
- Service Account: psibuilder@precision-sewer-inspection.iam.gserviceaccount.com
- Purpose: Booking schedule management

### AWS S3
- Purpose: Video/photo storage
- Presigned URLs for secure access

### Abacus Notifications
| Event | Recipient |
|-------|-----------|
| Contact Form Submission | Admin |
| Inspection Submitted | Admin |
| Video Ready | Client |
| Link Expiring | Client |
| Returned for Corrections | Technician |

## CONTACTS

| Role | Email |
|------|-------|
| Booking | booking@precisionsewerinspections.com |
| Support | support@precisionsewerinspections.com |
| Douglas | douglas@precisionsewerinspections.com |
| Ryan | ryan@precisionsewerinspections.com |

## DEVELOPMENT NOTES

### Key Considerations
1. Camera hardware integration (S830ASMKT) may enable live preview/control
2. AI-assisted report generation requires good training data
3. Voice transcription accuracy in noisy field environments
4. Offline resilience for areas with poor connectivity
5. Highlight reel algorithm for "key findings" selection

### Potential Gaps
- Offline mode implementation status unclear
- Push notifications (FCM) implementation status unclear
- Trip fee authorization workflow needs verification
- Multi-technician GPS map view in admin dashboard

### Code Quality
- Well-structured Next.js 14 App Router
- Comprehensive Prisma schema
- 15 specialized inspection stage components
- Good separation of concerns

## QUICK COMMANDS

```bash
# Start development
cd "/Volumes/Storage/PSI/precision_sewer_inspection (1)/nextjs_space"
yarn dev

# Database operations
npx prisma studio
npx prisma db push
npx prisma seed

# Build for production
yarn build
yarn start
```

## CACHED KNOWLEDGE KEYS

This skill synthesizes:
- `psi:project_knowledge_base`
- `psi:workflow_reference`
- `psi:technician_pwa_components`
- `psi:api_routes`
