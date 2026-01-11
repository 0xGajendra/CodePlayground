import type { Request, Response, NextFunction } from "express";

type RateLimitEntry = {
    count: number;
    windowStart: number;
};

const WINDOW_SIZE= 60000; // 1 minute
const MAX_REQUESTS= 5;

const store = new Map<string, RateLimitEntry>();

//Fixed Window Counter
export function rateLimiter(req: Request, res:Response, next: NextFunction){
    const ip: string | undefined = req.ip; //could be user id, apiKey anything that identifies a client as unique
    const currentTime = Date.now();

    const entry = store.get(ip as string);

    if(!entry){
        store.set(ip as string, {count: 1, windowStart: currentTime});
        return next();
    }
    if(currentTime - entry.windowStart > WINDOW_SIZE){
        store.set(ip as string, {count: 1, windowStart: currentTime}); //new widnow
        return next();
    }

    if(entry.count < MAX_REQUESTS){
        entry.count += 1;
        store.set(ip as string, entry);
        return next();
    }
    
    return res.status(429).send('Too many requests - try again later');

}

//Sliding Window Log
const store1 = new Map<string, number[]>();

export const slidingWindowLogRateLimiter = (req: Request, res: Response, next: NextFunction) =>{
    const ip: string = req.ip!;
    const currentTime = Date.now();

    const timestamps: number[] = store1.get(ip) || [];

    //remove request outside window frame
    const recent = timestamps.filter((time)=> currentTime-time <= WINDOW_SIZE);

    if(recent.length >= MAX_REQUESTS){
        return res.status(429).send('Too many requests - try again later');

    }

    recent.push(currentTime);
    store1.set(ip, recent);
    next();
}