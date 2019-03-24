import Link from 'next/link'
import Header from "../components/Header"
import MyLayout from "../components/MyLayout"
import React, { useState } from 'react';
import fetch from 'isomorphic-unfetch'

export default function Index() {
    const [value, setvalue] = useState(0);
    const [rating, setrating] = useState(0);
    const [count, setcount] = useState(0);

   return (
    <div>
      <MyLayout>
       <p>Faites votre recherche</p>
       <input onChange={(event) => setvalue(event.target.value)}/>
       <p><button onclick={postaddress(value, rating)}>Chercher</button></p>
       <p>Cette annonce a {count} pour une moyenne de {rating}</p>
      </MyLayout>
    </div>
  )
  console.log(rating)


  function postaddress(address, rating) {
   console.log("kk")
   fetch('http://localhost:8083', {
    method: 'POST',
    body: address,
   }).then(response => {
    return (response.json())
   }).then(data => {
    setrating(data.rating)
    setcount(data.count)
   });
  }
}
