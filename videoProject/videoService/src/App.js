import React from "react";
import axios from "axios";

class App extends React.Component{
  state = {
    isLoading : true
  };
  
  getMovies = async () => {
    console.log(await axios.get("localhost:9090"));
    this.setState({isLoading : false});
  }
  
  componentDidMount() {
    this.getMovies()  
  }

  render(){
    const {isLoading} = this.state;
    return (
      <div>
        {isLoading ? "Loading" : "asa"}
      </div>

    )
  }
}
export default App;
