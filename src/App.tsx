// import {GoogleMap,Marker,useJsApiLoader} from "@react-google-maps/api"
// import { useCallback, useState } from "react";

import { useEffect } from "react";

// function App() {
//   const {isLoaded} = useJsApiLoader({
//     id:"google-map-script",
//     googleMapsApiKey:""
//   })

//   const containerStyle = {
//     height: '100vh'
//   };

//   const center = {
//     lat: -6.2248764,
//     lng: 106.684606
//   };

//   const [map,setMap] = useState<google.maps.Map | null>(null)

//   const onLoad = useCallback((map:google.maps.Map)=>{
//     console.log("map data : ",map.data)
//     setMap(map)
//   },[])

//   const onUnMount = useCallback(()=>{
//     setMap(null)
//   },[])

//   return isLoaded ? (
//     <GoogleMap
//       mapContainerStyle={containerStyle}
//       center={center}
//       zoom={15}
//       onLoad={onLoad}
//       onUnmount={onUnMount}
//       mapContainerClassName="m-0 p-0"
//     >
//       <Marker position={center}></Marker>
//     </GoogleMap>
//   ) : <></>
// }

function App() {
  useEffect(() => {
    window.watch("1");
    window.onDataRecieved = (data: any) => {
      console.log("recieved data: ", data);
    };
  }, []);

  return <button>Click here to invoke WebAssembly!</button>;
}

export default App;
