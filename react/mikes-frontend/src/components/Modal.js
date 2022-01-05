function Modal(props) {
    return (
    <div className='modal'>
        <p>Are you sure?</p>
        <button className='btn btn--alt'>Confirm</button>
        <button className='btn'>Cancel</button>
    </div>
    );
}

export default Modal