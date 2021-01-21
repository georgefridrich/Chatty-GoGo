class App extends React.Component {
  render() {
    if (this.loggedIn) {
      return (<LoggedIn />);
    } else {
      return (<Home />);
    }
  }
}

class Home extends React.Component {
    render() {
      return (
        <div className="container">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>Chatty-GoGo</h1>
            <p>An exclusive message service for discerning users.</p>
            <p>Please register if this is your first visit.</p>
            <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
          </div>
        </div>
      )
    }
  }

class LoggedIn extends React.Component {
    constructor(props) {
    super(props);
    this.state = {
    userdata: []
    }
}
      
    render() {
      return (
        <div className="container">
          <div className="col-lg-12">
            <br />
            <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
            <h2>Chatty-GoGo</h2>
            <p>Where those who know, come to chat about it.</p>
          </div>
        </div>
      )
    }
  }

ReactDOM.render(<App />, document.getElementById('chatty'));