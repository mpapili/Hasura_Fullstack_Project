import { Link } from 'react-router-dom';

import  './MainNavigation.Module.css';

function MainNavigation() {

    return <header className='header'>
        <div className='logo'>React Meetups</div>
        <nav>
            <ul>
                <li>
                    <Link to='/'>All Meetups</Link>
                </li>
                <li>
                    <Link to='/new-meetup'>Add New Meetup</Link>
                </li>
                <li>
                    <Link to='/favorites'>MY Favorites</Link>
                </li>
            </ul>
        </nav>
    </header>

}

export default MainNavigation;