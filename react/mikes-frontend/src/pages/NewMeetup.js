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
     
    async function addMeetupHandler(meetupData) {
        const response = await fetch('http://127.0.0.1:8000/meetups/add',
        {   headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify(meetupData),
        });
    }

    async function addMeetupGraphql(meetupData) {
        const MUTATION = `
        mutation AddMeetup {
            insert_meetups_meetup_one(object: 
              {address: "${meetupData.address}", 
               description: "${meetupData.description}",
               image: "${meetupData.image}",
               title: "${meetupData.title}"}) {
              id
            }
          }          
        `
        const response = await fetch('http://localhost:8080/v1/graphql', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                query: MUTATION,
            }),
        })
        const responseBody = await response.json();
        console.log("Results of POST attempt:" + responseBody);
        
    }

    async function addMeetupGraphqlWithValidation(meetupData) {
        const MUTATION = `
        mutation AddMeetup {
            AddMeetup(meetupData: 
            {
                address: "${meetupData.address}",
                description: "${meetupData.description}",
                id: 10,
                title: "${meetupData.title}",
                image: "${meetupData.image}"
            })
            {
              address
              description
              id
              image
              title
            }
            insert_meetups_meetup_one(object:
            {
                address: "${meetupData.address}",
                description: "${meetupData.description}",
                id: 10,
                title: "${meetupData.title}",
                image: "${meetupData.image}"
            }) {
              id
            }
          }          
        `
        const response = await fetch('http://localhost:8080/v1/graphql', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                query: MUTATION,
            }),
        })
        const responseBody = await response.json();
        console.log("Results of POST attempt:" + responseBody);

    }

    return <section>
         <h1>Add New Meetup</h1>
        <NewMeetupForm onAddMeetup={addMeetupGraphql}/>
        <button onClick={addMeetupGraphqlWithValidation}>Test graphql</button>
    </section>

}

export default NewMeetupPage;