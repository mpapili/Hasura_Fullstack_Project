import { useState, useEffect } from "react";

import MeetupList from '../components/meetups/MeetupList';

/*
const DUMMY_DATA = [
    {
        id: 'm1',
        title: 'This is the first meetup!',
        image:
            'https://www.morristourism.org/wp-content/uploads/2017/05/Taylor-Ice-Cream-Weisler-Chester-scaled-2.jpg',
        address: 'MeetStreet 5, 12345 Super City',
        description: 'going to talk about how cool the town is',
    },
    {
        id: 'm2',
        title: 'This is the second meetup!',
        image:
            'https://upload.wikimedia.org/wikipedia/commons/thumb/f/fe/Main_Street%2C_Chester%2C_NJ_-_clock.jpg/1200px-Main_Street%2C_Chester%2C_NJ_-_clock.jpg',
        address: 'MeetStreet 7, 12345 Cool City',
        description: 'going to talk about how cool the town is again',
    },
]
*/


function AllMeetupsPage() {
    const [ allMeetups, setAllMeetups ] = useState([]);

    async function getMeetups() {
        const response = await fetch('http://localhost:8000/meetups'
        )
        .then(response => response.json())
        .then(data => setAllMeetups(data));
    }    

    useEffect(() => getMeetups(), []);

    return <section>
        <h1>All Meetups</h1>
        <MeetupList meetups={allMeetups} />
    </section>
}

export default AllMeetupsPage;