import { useState } from "react";

import NewMeetupForm from "../components/meetups/NewMeetupForm";


function NewMeetupPage() {

    async function addMeetupHandler(meetupData) {
        const response = await fetch('http://127.0.0.1:8000/meetups/add',
        {   headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify(meetupData),
        });
    }

    return <section>
         <h1>Add New Meetup</h1>
        <NewMeetupForm onAddMeetup={addMeetupHandler}/>
    </section>

}

export default NewMeetupPage;