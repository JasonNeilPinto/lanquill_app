import React from "react";
import { Link } from "react-router-dom";

const CyberAbout = () => {
  return (
    <section className="counter-with-video pt-80 ptb-120" id="cyber-about">
      <div className="container">
        <div className="row">
          <div className="col-lg-6 col-md-12">
            <div className="cyber-about-img text-center mb-30 mb-lg-0">
              <img src="/img/about_cyber.jpg" alt="VR" className="img-fluid" />
              <div className="row g-0">
                <div className="col-lg-5">
                  <div className="sheild-img">
                    <img
                      src="/img/about2.png"
                      alt="Sheild"
                      className="img-fluid d-none d-lg-block"
                    />
                  </div>
                </div>
                <div className="col-lg-6 col-md-12">
                  <div className="pe-2">
                    <div className="cyber-about-count-box d-md-flex bg-white p-4 mt-3">
                      <div className="pe-3">
                        <h2>170+</h2>
                      </div>
                      <div>
                        <h5 className="h6">Finished Projects</h5>
                        <p className="mb-0">Started in 2015</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="pt-5">
              <div className="section-heading mb-5">
                <h5 className="h6 text-primary">About Us</h5>
                <h2>Trusted by 50+ Happy Customers Worldwide</h2>
                <p>
                  Our technology and domain experts help you achieve success
                </p>
              </div>
              <div className="row">
                <div className="col-lg-6">
                  <div className="bg-white cyber-about-box mb-30 mb-lg-0">
                    <div className="cyber-about-icon">
                      <i className="far fa-hand-receiving"></i>
                    </div>
                    <h5 className="h-6">Who We Are</h5>
                    <p>
                      We are a multinational IT consulting and delivery services
                      organization, specializing in building innovative
                      solutions with new-age technologies.
                    </p>
                    <Link
                      to="/about-us"
                      className="text-decoration-none text-dark"
                    >
                      Explore more
                    </Link>
                  </div>
                </div>
                <div className="col-lg-6">
                  <div className="bg-white cyber-about-box">
                    <div className="cyber-about-icon">
                      <i className="far fa-users"></i>
                    </div>
                    <h5 className="h-6">Why choose us</h5>
                    <p>
                      We are agile, detail-oriented, and experienced in driving
                      digital and Gen AI transformation, cybersecurity, and
                      analytics.
                    </p>
                    <Link
                      to="/about-us"
                      className="text-decoration-none text-dark"
                    >
                      Explore more
                    </Link>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default CyberAbout;
