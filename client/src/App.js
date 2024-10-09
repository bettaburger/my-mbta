import { useState, useEffect } from 'react';
import axios from 'axios';

const App = () => {
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
        <ul>
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
              <li key={id}>
                <h2>{header}</h2>
                <p>Service Effect: {service_effect}</p>
                <p>Cause: {cause}</p>
                <p>Description: {description}</p>
                <p>Severity: {severity}</p>
                <p>Updated At: {new Date(updated_at).toLocaleString()}</p>
              </li>
            );
          })}
        </ul>
        <p>No alerts available.</p>
    </div>
  );
};

export default App;