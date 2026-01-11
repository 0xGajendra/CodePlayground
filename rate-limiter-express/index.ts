import express from 'express';
import { rateLimiter, slidingWindowLogRateLimiter } from './rateLimiter';
const  app = express();
const port = 3000;

app.use(express.json());
// app.use(rateLimiter);
app.use(slidingWindowLogRateLimiter)

app.get('/', (req, res)=>{
    res.send('Hello world!');
})

app.listen(port, ()=>{
    console.log(`Server is running at http://localhost:${port}`);
})