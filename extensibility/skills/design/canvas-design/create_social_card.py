#!/usr/bin/env python3
"""Create museum-quality GitHub social card for FLOYD Code."""

from PIL import Image, ImageDraw, ImageFont, ImageFilter
import os

# Paths
input_path = "original_workspace.png"
output_path = "floyd_social_card.png"
target_size = (1200, 630)

# Colors (RGBA)
PURPLE = (139, 92, 246, 255)
PINK = (236, 72, 153, 255)
PURPLE_SUBTLE = (139, 92, 246, 15)
BORDER = (168, 85, 247, 60)

# Load and resize image to cover target size
img = Image.open(input_path).convert("RGBA")
img_ratio = img.width / img.height
target_ratio = target_size[0] / target_size[1]

if img_ratio > target_ratio:
    # Image is wider - fit to height
    new_height = target_size[1]
    new_width = int(new_height * img_ratio)
else:
    # Image is taller - fit to width
    new_width = target_size[0]
    new_height = int(new_width / img_ratio)

img_resized = img.resize((new_width, new_height), Image.LANCZOS)

# Create canvas with target size
canvas = Image.new("RGBA", target_size, (15, 15, 20, 255))

# Center the resized image
paste_x = (target_size[0] - new_width) // 2
paste_y = (target_size[1] - new_height) // 2
canvas.paste(img_resized, (paste_x, paste_y))

# Create base with blur for depth (subtle vignette effect)
blurred = canvas.filter(ImageFilter.GaussianBlur(radius=25))
overlay = Image.new("RGBA", target_size, (0, 0, 0, 0))

# Add subtle purple-pink gradient overlay (bottom-up)
for y in range(target_size[1]):
    intensity = int((y / target_size[1]) * 20)  # Subtle at bottom
    if intensity > 0:
        overlay_alpha = intensity
        # Mix purple and pink based on position
        purple_amt = 1 - (y / target_size[1])
        pink_amt = y / target_size[1]
        r = int(PURPLE[0] * purple_amt + PINK[0] * pink_amt)
        g = int(PURPLE[1] * purple_amt + PINK[1] * pink_amt)
        b = int(PURPLE[2] * purple_amt + PINK[2] * pink_amt)
        for x in range(target_size[0]):
            overlay.putpixel((x, y), (r, g, b, overlay_alpha))

canvas = Image.alpha_composite(canvas, overlay)

# Add very subtle color wash
wash = Image.new("RGBA", target_size, PURPLE_SUBTLE)
canvas = Image.alpha_composite(canvas, wash)

# Draw refined border
draw = ImageDraw.Draw(canvas)
margin = 12
border_rect = [margin, margin, target_size[0] - margin, target_size[1] - margin]
# Outer glow effect (multiple thin lines)
for i in range(3):
    alpha = 30 - i * 8
    draw.rectangle(border_rect, outline=(168, 85, 247, alpha), width=1)
    border_rect[0] += 1
    border_rect[1] += 1
    border_rect[2] -= 1
    border_rect[3] -= 1

# Create text watermark
text_canvas = Image.new("RGBA", target_size, (0, 0, 0, 0))
text_draw = ImageDraw.Draw(text_canvas)

# Try to use system fonts, fallback to default
try:
    font_large = ImageFont.truetype("/System/Library/Fonts/Helvetica.ttc", 52)
    font_small = ImageFont.truetype("/System/Library/Fonts/Helvetica-Light.ttc", 22)
except:
    try:
        font_large = ImageFont.truetype("DejaVuSans-Bold.ttf", 52)
        font_small = ImageFont.truetype("DejaVuSans.ttf", 22)
    except:
        font_large = ImageFont.load_default(size=52)
        font_small = ImageFont.load_default(size=22)

# Position: bottom left with generous margin
text_x = 40
text_y = target_size[1] - 85

# "FLOYD" in white with subtle shadow
text_draw.text((text_x + 2, text_y + 2), "FLOYD",
               font=font_large, fill=(0, 0, 0, 80))
text_draw.text((text_x, text_y), "FLOYD",
               font=font_large, fill=(255, 255, 255, 245))

# "Code" in pink below, offset and smaller
text_draw.text((text_x + 3, text_y + 58), "Code",
               font=font_small, fill=(232, 121, 249, 220))

# Composite text onto canvas
canvas = Image.alpha_composite(canvas, text_canvas)

# Ensure output is pristine
canvas = canvas.convert("RGBA")
canvas.save(output_path, "PNG", optimize=True)

print(f"Created: {output_path}")
print(f"Size: {canvas.width}x{canvas.height}")
