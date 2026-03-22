# Skill: Ouroboros / Self-Evolution Protocol
🏷️ **Category:** core
💡 **Purpose:** Gives the LLM explicit instructions on how to use the "God Mode" OrbStack environment to clone its own repository, modify its own source code, test it inside a secure sandbox container, and optionally deploy the upgrades to the host globally.

---

## The Execution Pipeline

When the user asks you to "implement an upgrade to yourself" or invoke "Self-Evolution":

### Step 1: Boot the God-Mode Sandbox
Execute the `mcp_floyd-lab_spawn_lab` tool.
- Pass a unique `session_id`.
- The tool natively live-mounts the host's directory into `/workspace` and binds `/var/run/docker.sock`. YOU NOW HAVE ROOT DOCKER CONTROL from inside the VM.

### Step 2: Workspace Discovery & Modification
- Because the repository is already mounted at `/workspace`, DO NOT DO A `git clone`. The host's files are already there.
- Use your standard tools (`grep`, `glob`, `multiedit`) to locate the relevant Go or Node.js code that needs to be upgraded. Use the `manage_scratchpad` tool to outline your architectural plan.
- Execute structural refactors and verify the files match your planned changes.

### Step 3: Sandboxed Compilation & Verification
- Use `mcp_floyd-lab_execute_in_lab` to run:
  `go build ./...`
  `go test ./...`
- Because this occurs *inside the VM*, if the code panics or a dependency installation breaks, the user's host machine is completely insulated. Let the VM eat the panic.

### Step 4: Infrastructure Mutation (DooD)
- If your patch involves orchestrating other containers, use `mcp_floyd-lab_execute_in_lab` to run `docker run ...` or `docker build ...`. 
- Since the docker socket is mounted, the VM is commanding the host's OrbStack daemon directly.

### Step 5: The Overwrite (The Snake Eating its Tail)
If the tests pass 100% inside the VM sandbox and compilation produces zero exit codes:
- To apply the upgrade globally to the Host, run the standard build script via standard bash (Host context):
  ```bash
  ./scripts/build.sh
  rm /opt/homebrew/bin/floyd /opt/homebrew/bin/superfloyd
  cp floyd /opt/homebrew/bin/floyd
  cp superfloyd /opt/homebrew/bin/superfloyd
  ```

### Step 6: Confirmation
- Execute `/opt/homebrew/bin/floyd --version` to prove you successfully modified and re-deployed your own executable structure.
- Call `mcp_floyd-lab_teardown_lab` to clean up the sandbox.