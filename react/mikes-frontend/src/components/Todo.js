import { useState, useStates } from 'react';

import Modal from './Modal';
import Backdrop from './Backdrop';


function Todo(props) {
    // react hook! can only be called in function
    const [ modalIsOpen, setModalIsOpen] = useState(false); 

    function deleteHandler() {
        setModalIsOpen(true);
    }

    function closeModalHandler() {
        setModalIsOpen(false);
    }

    return (
        <div>
            <div className='card'>
                <h2>{props.text}</h2>
                <button className='btn' onClick={deleteHandler}>Delete</button>
            </div>
            {/* Render "Modal" if "modalIsOpen" is True */}
            {modalIsOpen && <Modal onCancel={closeModalHandler} onConfirm={closeModalHandler}/>}
            {/* Passing a function as a prop like an object!!! */}
            {modalIsOpen && <Backdrop onClick={closeModalHandler}/>}
        </div>
    )   
}

export default Todo;