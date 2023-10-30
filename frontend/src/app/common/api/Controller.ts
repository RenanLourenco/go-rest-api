import { Personality } from "@/app/types";
import axios, { AxiosInstance } from "axios"

export default class Controller {
    private api: AxiosInstance

    constructor(){
        this.api = axios.create({
            baseURL:'http://localhost:8000'
        })
    }

    async list():Promise<Personality[]> {
        const request = await this.api.get("/api/personalities");

        return request.data
    }

    async create(data : {name:string, history: string}): Promise<Personality> {
        const request = await this.api.post("/api/personalities",{
            name:data.name,
            history:data.history
        })

        return request.data
    }
}