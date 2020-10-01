import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="container">
      <div className="col-xl-10 col-lg-12 col-md-9 mx-auto">
        <div className="card o-hidden border-0 shadow-lg my-5">
          <div className="card-body p-0">
            <div className="row">
              <div className="col-lg-6 d-none d-lg-block">
                <img src="https://placehold.co/30x35" className="card-img"></img>
              </div>
              <div className="col-lg-6">
                <div className="p-5">
                  <div className="text-center">
                    <h1 className="h4 text-grey-900 mb-4">Hai, Welcome Back!</h1>
                  </div>
                  <form>
                    <div className="form-group mb-3">
                      <label> Email </label>
                      <input type="email" className="form-control" name="email"></input>
                    </div>
                    <div className="form-group mb-3">
                      <label> Password </label>
                      <input type="password" className="form-control" name="password"></input>
                    </div>
                    <button className="btn btn-primary btn-block" type="submit">Login</button>
                  </form>
                  <hr></hr>
                  <div className="text-center">
                    <a href="#" className="small">Forgot Password ?</a>
                  </div>
                  <div className="text-center">
                    <a href="#" className="small">Create Your Account !</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
