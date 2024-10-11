import React from 'react'
import button from '../styles/button.css'
function Nav() {
  return (
    <div class="navbar">
        <a href="#Alerts">Alerts</a>
        <div class="dropdown">
            <button class="dropbtn">Subway
            <i class="fa fa-caret-down"></i>
            </button>
            <div class="dropdown-content">
            <a href="#">Red Line</a>
            <a href="#">Orange Line</a>
            <a href="#">Green Line</a>
            <a href="#">Blue Line</a>
        </div>
    </div> 
    </div>
  )
}

export default Nav