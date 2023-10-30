
import Controller from '@/app/common/api/Controller';
import { Personality } from '@/app/types';
import Table from '@mui/joy/Table';
import axios from 'axios';
import AddPersonality from '../AddPersonality/AddPersonality';
import DeletePersonality from '../DeletePersonality/DeletePersonality';



export default async function PersonalitiesContent() {
    const controller = new Controller()
    const personalities = await controller.list()

    return (
        <div>
            <Table variant="soft" >
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>History</th>
                    </tr>
                </thead>
                <tbody>
                    {personalities.map((personality) => {
                        return <tr key={personality.id}>
                            <td >{personality.name}</td>
                            <td>{personality.history}</td>
                            <td><DeletePersonality id={personality.id}/></td>
                        </tr>
                    })}
                </tbody>
            </Table>
            <AddPersonality/>
        </div>
    )
}