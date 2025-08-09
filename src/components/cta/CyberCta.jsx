import React from "react";
import { Link } from "react-router-dom";

const CyberCta = () => {
  return (
    <section className="cyber-cta pb-60">
      <div className="container">
        <div
          className="cyber-cta-bg p-5 rounded-custom"
          style={{
            background:
              "url('/img/cb_call_to_action.png') no-repeat center / cover",
          }}
        >
          <div className="row">
            <div className="col-lg-9 col-md-9">
              <div className="cyber-cta-info position-relative">
                <div className="">
                  <h2 className="text-white">
                    Are You Ready? Book Appoinment Now!
                  </h2>
                  <p className="lead text-white mb-0">Call : +91 9620555571</p>
                </div>
              </div>
            </div>
            <div className="col-lg-3">
              <div className="cyber-cta-btn">
                <Link
                  to="/request-demo"
                  className="mt-3 btn btn-primary me-auto"
                >
                  Book a Demo <i className="far fa-arrow-right"></i>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default CyberCta;
