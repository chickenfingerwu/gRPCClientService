import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import clientsList from './client'

const App = () => {
  const [clients, setClients] = useState([])
  useEffect(() => {
        clientsList
            .getAll()
            .then(initialClients => {
              setClients(initialClients)
            })
      },[]);

  const renderClients = () => {
      let names = [];
      clients.forEach(function (item, index, names) {
          names.push(item.name)
      })
      return names
  }
  return(
      <div>
        <h2>Clients</h2>
      </div>
  )
}

export default App;
