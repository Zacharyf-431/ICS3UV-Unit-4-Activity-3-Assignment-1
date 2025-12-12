/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-12-10
 * @fileoverview: This file calculates total of items and price in a shopping cart
 */



// assign constants 
const DISCOUNT_THRESHOLD = 350;
const DISCOUNT_RATE = 0.1;
const HST_RATE = 0.13;

// assign variables
let numItems: number;
let subtotal: number = 0;
let discount: number = 0;
let hst: number;
let total: number;

// assign arrays
let itemNames: string[];
let itemPrices: number[];

// INPUT
numItems = Number(prompt("Enter the number of items you are going to purchase: ")) || 0;
itemNames = new Array(numItems);
itemPrices = new Array(numItems);

for (let i = 0; i < numItems; i++) {
  itemNames[i] = prompt(`Enter the name of item ${i + 1}:`) || "";
  itemPrices[i] = Number(prompt(`Enter the price of ${itemNames[i]}:`)) || 0;
  subtotal += itemPrices[i];
}

// PROCESS
if (subtotal > DISCOUNT_THRESHOLD) {
  discount = subtotal * DISCOUNT_RATE;
}
subtotal -= discount;
hst = subtotal * HST_RATE;
total = subtotal + hst;

// OUTPUT
console.log(`Your shopping cart includes: ${itemNames.join(", ")}`);
console.log(`The total amount of items in your shopping cart is ${numItems}`);

console.log(`The subtotal cost of your shopping trip was $${(subtotal + discount).toFixed(2)}`);

if (discount > 0) {
  console.log(`You are eligible for a discount of $${discount.toFixed(2)}`);
}

console.log(`The HST is $${hst.toFixed(2)}`);
console.log(`The total is $${total.toFixed(2)}`);

console.log("\nDone.");