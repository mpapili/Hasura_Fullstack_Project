import axios from 'axios';
import { useState } from "react";

import NewMeetupForm from "../components/meetups/NewMeetupForm";


function NewMeetupPage() {

    // Test with the sample "Pets" endpoint I'd created earlier
    async function getRequestTest() {
        const response = await fetch('http://127.0.0.1:8000/pets/1'
        )
        .then(response => response.text())
        .then(data => setTestData(data));
    }

    const [ testData, setTestData ] = useState(''); 

    async function addMeetupHandler(meetupData) {
        console.log(meetupData);
        const response = await fetch('http://127.0.0.1:8000/meetups/add',
        {   headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify(meetupData),
        });
    }

    return <section>
         <h1>Add New Meetup</h1>
        <NewMeetupForm onAddMeetup={addMeetupHandler}/>
        <button onClick={getRequestTest}>Test Request</button>
        {testData}
    </section>

}

export default NewMeetupPage;