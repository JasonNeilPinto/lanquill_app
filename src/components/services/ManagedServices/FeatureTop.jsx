import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">Managed Services</h2>
              <p className="mb-4">
                We deliver cost-efficient, scalable, and high-quality IT Managed
                Services built on industry-standard ITIL service management
                frameworks. Our approach ensures consistent service delivery,
                faster response times, and improved user satisfaction—allowing
                your business to focus on growth while we handle operational
                complexity.
              </p>
              <p className="mb-4">
                We offer comprehensive managed services that are proactive,
                secure, and aligned with your digital goals.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/ManagedServices.png"
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
