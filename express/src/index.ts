/**
 * Required External Modules
 */
import * as dotenv from "dotenv";
import express from "express";
import cors from "cors";
import helmet from "helmet";

import {itemsRouter} from './items/items.router'
dotenv.config(); // get our environment variables!

/**
 * App Variables
 */

// did we get a port from the ".env" file?
if (!process.env.PORT) {
    console.log("no port!");
    process.exit(1);
 }
 const PORT: number = parseInt(process.env.PORT as string, 10);

 const app = express();
 
/**
 *  App Configuration
 */
app.use(helmet());
app.use(cors());
app.use(express.json());

app.use("/api/menu/items", itemsRouter);


/**
 * Server Activation
 */
// fire it off!
app.listen(PORT, () => {
    console.log(`Listening on port ${PORT}`);
})
