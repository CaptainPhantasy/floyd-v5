---
name: csv-data-summarizer
description: Analyzes CSV files, generates summary stats, and plots quick visualizations using Python and pandas. Automatically triggers when user uploads or references a CSV file, asks to summarize/analyze/visualize tabular data, or wants insights from CSV data.
---

# CSV Data Summarizer

This skill analyzes CSV files and provides comprehensive summaries with statistical insights and visualizations.

## CRITICAL BEHAVIOR REQUIREMENT

Do not ask the user what they want. Do not offer options. Do not wait.

Immediately:
1. Load the CSV into pandas
2. Inspect structure (columns, types, missingness)
3. Detect dataset type (sales, customer, financial, operational, survey, generic)
4. Run all relevant analyses for that type
5. Generate all relevant visualizations
6. Present everything in one complete analysis

## Workflow

### 1. Load and Inspect

```python
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

# Load CSV
df = pd.read_csv(file_path)

# Inspect structure
print(f"Shape: {df.shape}")
print(f"\nColumns: {df.columns.tolist()}")
print(f"\nData types:\n{df.dtypes}")
print(f"\nMissing values:\n{df.isnull().sum()}")
print(f"\nFirst few rows:\n{df.head()}")
```

### 2. Detect Dataset Type

Based on column names and content, classify as:
- **Sales**: revenue, amount, price, quantity, product
- **Customer**: customer_id, name, email, phone, address
- **Financial**: balance, transaction, debit, credit, account
- **Operational**: timestamp, status, duration, user_id, action
- **Survey**: rating, score, response, question
- **Generic**: doesn't match patterns above

### 3. Type-Specific Analyses

#### Sales Data
- Revenue by product/category
- Sales trends over time
- Top products by volume/revenue
- Average order value

#### Customer Data
- Customer demographics
- Geographic distribution
- Signup trends over time
- Churn analysis (if dates present)

#### Financial Data
- Transaction volumes
- Balance distributions
- Debit vs credit analysis
- Account activity patterns

#### Operational Data
- Status breakdown
- Duration distributions
- Peak usage times
- User activity patterns

#### Survey Data
- Response distributions
- Average scores by category
- Rating trends
- Sentiment analysis (if text present)

#### Generic Data
- Numeric column statistics (mean, median, std, min, max)
- Categorical value counts
- Correlation analysis
- Distribution plots

### 4. Generate Visualizations

Create 3-5 relevant plots:
- Time series (if date columns present)
- Distributions (histograms, box plots)
- Categorical breakdowns (bar charts)
- Correlations (heatmap if numeric columns)
- Scatter plots (for relationships)

```python
# Example: Time series
if 'date' in df.columns:
    df['date'] = pd.to_datetime(df['date'])
    df.groupby('date')['revenue'].sum().plot()
    plt.title('Revenue Over Time')
    plt.show()

# Example: Distribution
if 'amount' in df.columns:
    df['amount'].hist(bins=30)
    plt.title('Amount Distribution')
    plt.show()

# Example: Category breakdown
if 'category' in df.columns:
    df['category'].value_counts().plot(kind='bar')
    plt.title('Count by Category')
    plt.show()
```

### 5. Present Results

Format: Present in one complete message with:
- Dataset overview (rows, columns, date range if applicable)
- Key findings (2-3 bullets)
- Statistical summary
- All visualizations
- Data quality notes (missing values, outliers)

## Forbidden Phrases

Never say:
- "What would you like to do with this data?"
- "Would you like me to..."
- "I can show you..."
- "Let me know if you want..."
- "What kind of analysis are you interested in?"

Always:
- Just do the analysis
- Present complete results
- Be proactive about what's interesting

## Implementation Notes

```python
def summarize_csv(file_path):
    """
    Complete CSV analysis pipeline
    """
    # 1. Load
    df = pd.read_csv(file_path)
    
    # 2. Inspect
    shape = df.shape
    columns = df.columns.tolist()
    dtypes = df.dtypes
    missing = df.isnull().sum()
    
    # 3. Detect type
    dataset_type = detect_type(df)
    
    # 4. Run type-specific analysis
    analysis = analyze_by_type(df, dataset_type)
    
    # 5. Generate plots
    plots = create_visualizations(df, dataset_type)
    
    # 6. Present
    print_summary(shape, columns, dtypes, missing, analysis, plots)
```

## Requirements

Dependencies:
- python >= 3.8
- pandas >= 2.0.0
- matplotlib >= 3.7.0
- seaborn >= 0.12.0
