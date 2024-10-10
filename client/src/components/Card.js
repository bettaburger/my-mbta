import { useState, useEffect } from 'react';
import axios from 'axios';
import card from '../styles/card.css'


const Alerts = () => {
    const [alerts, setAlerts] = useState([]);

   // Fetching http request to port 9000 server
   useEffect(() => {
      axios
         .get("http://localhost:9000/alerts")
         .then((response) => {
            setAlerts(response.data.data);
         })
         .catch((err) => {
            console.log(err);
         });
   },[]);

    return (
        <div>
          <h1>MBTA Alerts</h1>
            <div>
              {alerts.map((alert) => { //Render the data out as a list. 
                const {
                id,
                attributes: {
                    header,
                    service_effect,
                    cause,
                    description,
                    severity,
                    updated_at
                  }
                } = alert;
    
                // Displaying the data alert fields.
                return (
                <div key={id}>
                    <div class='card'>
                    <h1 style={{color: "red"}}>{header}</h1>
                    <p>Service Effect: {service_effect}</p>
                    <p>Cause: {cause}</p>
                    <p>Description: {description}</p>
                    <p>Severity: {severity}</p>
                    <p>Updated At: {new Date(updated_at).toLocaleString()}</p>
                    </div>
                </div>
                );
              })}
            </div>
        </div>
      );
    };    

export default Alerts