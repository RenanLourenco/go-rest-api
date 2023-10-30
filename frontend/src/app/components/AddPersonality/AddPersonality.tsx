"use client"
import { Button, Modal, ModalClose, ModalDialog, Typography, Textarea } from "@mui/joy";
import React, { FormEvent, useState } from "react";
import Input from '@mui/joy/Input';
import Controller from "@/app/common/api/Controller";





export default function AddPersonality() {
    const controller = new Controller()

    const [modalOpen, setModalOpen] = useState(Boolean)
    const [name, setName] = useState('')
    const [history, setHistory] = useState('')
    
    const addPersonality = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const createdPersonality = await controller.create({name, history})
        window.location.reload();
    }

    return (
        <div className="mt-4">
            <button onClick={() => setModalOpen(true)} className="bg-lime-500 rounded-lg p-3 text-white">Add personality</button>
            <Modal open={modalOpen} onClose={() => setModalOpen(false)}>
                <ModalDialog>
                 <ModalClose/>
                 <Typography level="h4">Insert personality</Typography>
                 <form onSubmit={addPersonality}>
                    <Input placeholder="Name" required className="mb-2" value={name} onChange={(e) => setName(e.target.value)}/>
                    <Textarea placeholder="History" required size="lg" value={history} onChange={(e) => setHistory(e.target.value)}/>
                    <button  className="bg-blue-400 rounded-lg p-2 mt-3 text-white" type="submit">Submit</button>
                 </form>
                </ModalDialog>
            </Modal>
        </div>
        
    )


}