import React from "react";
import { Link } from "react-router-dom";

import SectionTitle from "../common/SectionTitle";

const CtaTwo = () => {
  return (
    <>
      <section className="cta-subscribe bg-dark ptb-80 position-relative overflow-hidden">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-lg-8 col-md-10">
              <div className="subscribe-info-wrap text-center position-relative z-2">
                <SectionTitle
                  subtitle="Let's Get Started!"
                  title="Elevate Your Business Today"
                  description="Experience how we help businesses like yours achieve remarkable success."
                  dark
                />
                <div className="form-block-banner mw-60 m-auto mt-5">
                  <Link to="/contact-us" className="btn btn-primary">
                    Contact with Us
                  </Link>
                </div>
              </div>
            </div>
          </div>
          <div
            className="
              bg-circle
              rounded-circle
              circle-shape-3
              position-absolute
              bg-blue
              left-5
            "
          ></div>
          <div
            className="
              bg-circle
              rounded-circle
              circle-shape-1
              position-absolute
              bg-green
              right-5
            "
          ></div>
        </div>
      </section>
    </>
  );
};

export default CtaTwo;
