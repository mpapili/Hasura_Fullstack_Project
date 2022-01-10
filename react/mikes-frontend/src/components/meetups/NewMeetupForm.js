import { useRef } from 'react';

import Card from "../ui/Card";
import classes from './NewMeetupForm.module.css';

function NewMeetupForm(props) {

    // Form that will be used to add NEW meetups
    // once again we'll wrap this in a "card" component

    // "onSubmit" auto-passes in an "event" object!
    // becomes an "onSubmit" for the FORM object

    // connected to the "title" input in the JSX
    const titleInputRef = useRef();
    const imageInputRef = useRef();
    const addressInputRef = useRef();
    const descriptionInputRef = useRef();


    function submitHandler(event) {
        event.preventDefault(); // stop default browser action of submitting to server

        // read the values
        const enteredTitle = titleInputRef.current.value;
        const enteredImage = imageInputRef.current.value;
        const enteredAddress = addressInputRef.current.value;
        const enteredDescription = descriptionInputRef.current.value;

        // create a map 
        const meetupData = {
            title: enteredTitle,
            image: enteredImage,
            address: enteredAddress,
            description: enteredDescription,
        }

        console.log("Our Data is now " + meetupData);

        // send our meetup data to a server!
        // we'll want to pass this up to our parent component to handle via
        // -a prop function that was passed in through "props"
        props.onAddMeetup(meetupData);
    }

    return <Card>
        <form className={classes.form} onSubmit={submitHandler}>
            <div className={classes.control}>
                {/*Note the Syntax! "htmlFor" in <label> matches "id" in <input>! */ }
                <label htmlFor="title">Meetup Title</label>
                <input type="text" required id="title" ref={titleInputRef}/>

                <label htmlFor="image">Image</label>
                <input type="url" required id="image" ref={imageInputRef}/>

                <label htmlFor="address">Address</label>
                <input type="text" required id="address" ref={addressInputRef} />

                <label htmlFor="description">Description</label>
                <textarea type="text" required id="description" rows='5' ref={descriptionInputRef}/>

            </div>
            <div className={classes.actions}>
                <button>Add Meetup</button>
            </div>
        </form>
    </Card>
}

export default NewMeetupForm;