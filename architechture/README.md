# 🍽️ Restaurant Order Management System

This architecture represents how a restaurant manages menus, food items, customer orders, tables, and invoices. It models the real-world process of ordering food, itemizing it, and generating a bill.

---

## 📌 **Major Entities in the Architecture**

1. **Menu**
2. **Food Item**
3. **Order Item**
4. **Order**
5. **Table**
6. **Invoice**

Each entity is connected to others in a way that fully represents a restaurant’s operational workflow.

---

## 1️⃣ Menu → Food Item

### **Menu**
The Menu table stores food categories such as:

- Starters  
- Main Course  
- Desserts  
- Beverages  

**Fields:**
- `ID`
- `Name`
- `Category`

### **Relationship**
A Menu category contains multiple **Food Items**, meaning one menu → many foods.

---

## 2️⃣ Food Item (Food)

Food items are the actual dishes available for customers to order.

**Fields:**
- `ID`
- `Name`
- `Price`

### **Relationship**
**Menu → Food Item**  
Each food item belongs to a category in the menu.

---

## 3️⃣ Order Item

Order Items represent each individual dish ordered by the customer.

**Fields:**
- `ID`
- `Qty` (Quantity)
- `Unit Price`

### **Relationships**
- **Order Item → Food Item**  
  Each order item corresponds to a specific food item.
- **Order Item → Order**  
  Multiple order items together form a complete order.

---

## 4️⃣ Order

An Order represents the complete request made by a customer.

**Fields:**
- `ID`
- `Order Date`

### **Relationships**
- **Table → Order**  
  A table can have multiple orders over time.
- **Order → Order Item**  
  An order contains multiple order items.

---

## 5️⃣ Table

Represents the restaurant tables where customers are seated.

**Fields:**
- `ID`
- `Table No.`
- `No of Guests`

### **Relationship**
**Table → Order**  
Each order is associated with a specific table.

---

## 6️⃣ Invoice

Invoices are generated after an order is completed.

**Fields:**
- `ID`
- `Payment Method`
- `Payment Status` (e.g., Paid, Pending)

### **Relationship**
**Invoice → Order**  
Each invoice is linked to a specific customer order.

---

## 🧩 Complete System Flow (Summary)

1. Menu defines the categories of food.
2. Food Items belong to each menu category.
3. A customer sits at a table and gives an order.
4. The Order record is created under that table.
5. Each food item inside the order is stored as an Order Item.
6. After completion, an Invoice is generated for billing.

---

## 🎯 Why This Architecture Works Well

- **Normalized structure** → No unnecessary duplicate data.  
- **Scalable design** → Can easily add waiters, discounts, GST, etc.  
- **Clear relationships**:
  - Menu → Food  
  - Order → Order Items  
  - Order → Invoice  
  - Table → Order  

---

## 🎉 Simple Real-World Analogy

- **Menu** → List of categories  
- **Food** → Actual dishes  
- **Table** → Where the customer sits  
- **Order** → What the customer ordered  
- **Order Item** → Individual dishes inside the order  
- **Invoice** → Final bill  

---
