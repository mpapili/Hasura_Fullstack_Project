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
            'https://www.morristourism.org/wp-content/uploads/2017/05/Taylor-Ice-Cream-Weisler-Chester-scaled-2.jpg',
        address: 'MeetStreet 7, 12345 Cool City',
        description: 'going to talk about how cool the town is again',
    },

]

function AllMeetupsPage() {

    return <section>
        <h1>All Meetups</h1>
        <ul>
            {DUMMY_DATA.map((meetup) => {
                return <li key={meetup.id}>{meetup.title}</li>
            })}
        </ul>
    </section>
}

export default AllMeetupsPage;