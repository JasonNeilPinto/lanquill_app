import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">Cybersecurity Services</h2>
              <p className="mb-4">
                Our comprehensive cybersecurity services are designed to protect
                your infrastructure, data, and users from emerging threats. We
                provide end-to-end solutions spanning cloud protection, data
                security, identity management, compliance, and advanced AI risk
                mitigation.
              </p>
              <p className="mb-4">
                Whether you are modernizing your tech stack, scaling securely in
                the cloud, or integrating AI systems, we ensure your
                organization remains secure, compliant, and resilient—every step
                of the way.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/Cybersecurity.png"
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
