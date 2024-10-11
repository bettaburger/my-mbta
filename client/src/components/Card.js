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
          <div className="card">
              {alerts.map((alert) => { //Render the data out as a list. 
                const {
                id,
                attributes: {
                    service_effect, // -> also use for filtering, bus, rail, access, etc.
                    header,
                    cause,
                    description,
                    //severity, -> use later for filtering, severity 1-10
                    updated_at
                  }
                } = alert;

                // Displaying the data alert fields.
                return (
                <div key={id} className="card-item">
                    <h1>{service_effect}</h1>
                    <h2>{header}</h2>
                    <p>Cause: {cause}</p>
                    <p>Description: {description}</p>
                    <p>Updated At: {new Date(updated_at).toLocaleString()}</p>
                    </div>
                );
              })}
            </div>
        </div>
      );
    };    

export default Alerts