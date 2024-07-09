- How to remove extra spaces from text in excel 
ctrl-A and then ctrl-H Opens Find and Replace
Find: space
Replace: No space
OK



Excel: Progress bar

=REPT(text, number, )
= REPT("|", C3) Change font to PlayBill for continous bar


𝟲 𝗪𝗮𝘆𝘀 𝘁𝗼 𝗙𝗶𝗻𝗱 𝗗𝗶𝘀𝘁𝗶𝗻𝗰𝘁 𝗩𝗮𝗹𝘂𝗲𝘀 𝗶𝗻 𝗘𝘅𝗰𝗲𝗹

- How to find distinct values in Excel using six different methods!
- UNIQUE function to array formulas

```
=ARRAY_CONSTRAIN(ARRAYFORMULA(UNIQUE('DISTINCT ITEMS'!$A$4:$A$28)), 14, 1)
```  
```
=IFERROR(INDEX($A$4:$A$28, MATCH(0,INDEX(COUNTIF(K$3:K3, $A$4:$A$28),0,0),0)),"")
```


```
𝗥𝗲𝗴𝘂𝗹𝗮𝗿 𝗙𝗼𝗿𝗺𝘂𝗹𝗮:
=IFERROR(INDEX($A$4:$A$28, MATCH(0,INDEX(COUNTIF(K$3:K3, $A$4:$A$28),0,0),0)),"")
𝗔𝗿𝗿𝗮𝘆 𝗙𝗼𝗿𝗺𝘂𝗹𝗮:
=IFERROR(INDEX($A$4:$A$28,MATCH(0,COUNTIF($M$3:M3,$A$4:$A$28),0)),"")

```

```
=VSTACK(FILTER(A2:G11, D2:D11="North"))

=VSTACK(FILTER(A2:G11, (C2:C11="Laptop") + (C2:C11="Tablet")))

=VSTACK(FILTER(A2:G11, DAY(B2:B11) <= 15))

=VSTACK(FILTER(A2:G11, C2:C11 <> "Laptop"))

=VSTACK(FILTER(A2:G11, MONTH(B2:B11) = 1))

```

- Excel: 6 Tricks of VSTACK with FILTER Function
- How to combine the VSTACK and FILTER functions to manipulate and analyze your sales data in advanced ways. 
- Can filter and aggregate data based on various conditions, allowing for more sophisticated data analysis and reporting tasks.
- FILTER selects the relevant rows, and VSTACK stacks them into a single dataset.
- FILTER selects the relevant rows, and 
- VSTACK stacks them into a single dataset.

1. Focuses on specific regions.
2. Multi-product data aggregation.
3.  isolates a specific time period.
4.  Excludes specific data points.
5.  Focuses on a specific month.
6. Consolidates product data across regions.

1. Combine Sales Data for North Region
Formula: =VSTACK(FILTER(A2:G11, D2:D11="North"))
Use this formula to filter and combine all sales data where the Region is "North". 
FILTER selects the relevant rows, and 
VSTACK stacks them into a single dataset. 
Great for analyzing performance metrics specific to the North region.

2. Combine Sales Data for Products Laptop and Tablet
Formula: =VSTACK(FILTER(A2:G11, (C2:C11="Laptop") + (C2:C11="Tablet")))
This formula filters data to include only "Laptop" and "Tablet" products. 
FILTER uses a logical OR condition to select rows for these products, and 
VSTACK combines them. 
It’s useful for comparing or aggregating sales data across multiple products.

3. Combine Data for Sales in the First Half of January
Formula: =VSTACK(FILTER(A2:G11, DAY(B2:B11) <= 15))
This filters sales data to include only transactions from the first 15 days of January. 
DAY(B2:B11) <= 15 identifies these days, and VSTACK combines the filtered rows. 
It’s helpful for early-month performance assessments.

4. Combine Data for All Products Except Laptop
Formula: =VSTACK(FILTER(A2:G11, C2:C11 <> "Laptop"))
Here, the formula excludes rows where the Product is "Laptop". 
FILTER applies the condition <> "Laptop", and 
VSTACK consolidates the remaining data. U
se this to focus on products other than "Laptop".

5. Stack Data for All Dates in January
Formula: =VSTACK(FILTER(A2:G11, MONTH(B2:B11) = 1))
This formula filters to include only data from January. 
MONTH(B2:B11) = 1 checks for the month, and 
VSTACK stacks the results. 
It’s ideal for gathering and analyzing data for a specific month.


6. Combine Sales Data from the Same Product Across Regions
Formula: =VSTACK(FILTER(A2:G11, C2:C11="Phone"))
Use this formula to filter and combine data where the Product is "Phone". 
FILTER selects all "Phone" sales, and VSTACK combines them. 
This  is perfect for analyzing the performance of a single product across different regions.
