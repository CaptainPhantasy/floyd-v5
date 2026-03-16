#!/usr/bin/env python3
"""Test Codebase_Cartographer application in headed mode."""

from playwright.sync_api import sync_playwright

with sync_playwright() as p:
    # Launch browser in headed mode
    browser = p.chromium.launch(
        headless=False,
        args=['--start-maximized']
    )
    context = browser.new_context(viewport={'width': 1920, 'height': 1080})
    page = context.new_page()

    print("Navigating to http://localhost:3000 ...")
    page.goto('http://localhost:3000')

    # Wait for page to fully load
    page.wait_for_load_state('networkidle')
    print("Page loaded - waiting for app to initialize...")

    # Wait a bit for React to render
    page.wait_for_timeout(2000)

    # Check the page title
    title = page.title()
    print(f"Page Title: {title}")

    # Take a screenshot
    screenshot_path = '/tmp/codebase_cartographer_screenshot.png'
    page.screenshot(path=screenshot_path, full_page=True)
    print(f"Screenshot saved to: {screenshot_path}")

    # Inspect the page content
    body_text = page.locator('body').text_content()
    print(f"\nPage content preview (first 500 chars):")
    print(body_text[:500] if body_text else "No content found")

    # Check for React app mounting
    root = page.locator('#root')
    root_count = root.count()
    print(f"\n#root element found: {root_count > 0}")

    # List all buttons on the page
    buttons = page.locator('button').all()
    print(f"\nButtons found: {len(buttons)}")
    for i, btn in enumerate(buttons[:10]):  # First 10 buttons
        print(f"  [{i}] {btn.text_content() or '<empty>'}")

    # List all inputs on the page
    inputs = page.locator('input').all()
    print(f"\nInputs found: {len(inputs)}")
    for i, inp in enumerate(inputs[:10]):
        inp_type = inp.get_attribute('type') or 'text'
        inp_placeholder = inp.get_attribute('placeholder') or ''
        print(f"  [{i}] type={inp_type}, placeholder={inp_placeholder}")

    # Keep browser open for visual inspection
    print("\n" + "="*50)
    print("Browser is open for visual inspection.")
    print("Press Enter in the terminal to close the browser...")
    print("="*50)

    input()  # Wait for user to press Enter

    browser.close()
    print("Browser closed.")
