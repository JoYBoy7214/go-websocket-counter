import React from 'react'
import ReactDOM from 'react-dom/client'
import { useEffect, useState } from 'react'

import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const [socket,setSocket]=useState(null)
   const ws=new WebSocket('ws://localhost:8080/ws')
  useEffect(()=>{
   
    ws.onopen=()=>{
      console.log("websocket connection established")
    }
    ws.onerror=(e)=>{
      console.warn("error :",e)
    }
    ws.onmessage=(e)=>{
      try {
        const message=JSON.parse(e.data)
        if(message.operation==="+"){
          
          setCount(prev=>prev+1)
        }
        else{
          setCount(prev=>prev-1)
        }
      } catch (error) {
        console.warn("error in parsing ",error)
      }
    }
    setSocket(ws)
  },[])

  const handleDecrement=()=>{
     
     if (socket && socket.readyState === WebSocket.OPEN){
     var message={operation:"-"}
        message=JSON.stringify(message)
          socket.send(message)
     }else{
      console.log('WebSocket is not open.');
     }
  };
  const handleIncrement=()=>{
    
     if (socket && socket.readyState === WebSocket.OPEN){
     var message={operation:"+"}
      message=JSON.stringify(message)
          socket.send(message)
     }
     else{
      console.log('WebSocket is not open.');
     }
  };

  return (
    <>
      <div >
        <button className='button' onClick={
          handleDecrement
        }>Decrementer</button>
        <button className="button">
          count is {count}
        </button>
        <button className='button' onClick={
           handleIncrement
        }>Incrementer</button>
      </div>
    </>
  )
}

export default App
