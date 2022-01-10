import NewMeetupForm from "../components/meetups/NewMeetupForm";
import { useQuery } from 'react-query';

const MEETUPS_QUERY = `
{
    meetups_meetup {
        address
        description
        id
        image
        title
      }    
}
`;

function NewMeetupPage() {

    async function testGraphql() {
        const response = await fetch('http://localhost:8080/v1/graphql', {
           method:'POST',
     
           headers:{'content-type':'application/json'},
           body:JSON.stringify({query: MEETUPS_QUERY})
        })
     
        const rsponseBody = await response.json();
        return rsponseBody.data;
     
        console.log("end of function")
     }
     
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
        <button onClick={testGraphql}>Test graphql</button>
    </section>

}

export default NewMeetupPage;