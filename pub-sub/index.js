import express from 'express';
import { eventBus } from './eventBus.js';
import morgan from 'morgan';

const app = express();
const PORT  =3000;

app.use(express.json());
app.use(morgan('dev'));

function saveOrder(order) {
    // Logic to save order
    console.log('Order saved:', order);
    return order ? order : null;
}

app.post('/order', (req, res) =>{
    const {order} = req.body; 
    const savedOrder = saveOrder(order);// Assume saveOrder is a function that saves the order
    if (savedOrder) {
        eventBus.emit('orderCreated', savedOrder);
        res.status(201).send({message: 'Order created successfully', order: savedOrder});
    } else {
        res.status(500).send({message: 'Failed to create order'});
    }
})

app.get('/', (req, res)=>{
    res.send('Welcome Pub_sub')
})

app.listen(PORT, ()=>{
    console.log(`Server is running on http://localhost:${PORT}`);
})