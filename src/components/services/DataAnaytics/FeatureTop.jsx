import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">
                Data and Analytics Services and Solutions
              </h2>
              <p className="mb-4">
                We help enterprises turn raw data into actionable insights
                through modern engineering, marketing analytics, and real-time
                reporting. With a focus on legacy modernization and governance,
                our future-ready solutions drive smarter decisions and scalable
                growth.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/DataAnalytics.png"
                  className="img-fluid"
                  alt="Mocup"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
