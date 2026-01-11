// eventBus.js
import EventEmitter from "events"
export const eventBus = new EventEmitter()

eventBus.on('orderCreated', async(order)=>{
    await SendEmail(order);
    await updateAnalytics(order);
})

async function SendEmail(order) {
    // Simulate email sending logic
    console.log(`Sending email for order: ${JSON.stringify(order)}`);
}

async function updateAnalytics(order) {
    // Simulate analytics update logic
    console.log(`Updating analytics for order: ${JSON.stringify(order)}`);
}