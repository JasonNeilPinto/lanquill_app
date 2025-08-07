import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">Security Testing Services</h2>
              <p className="mb-4">
                Our Security Testing Services are designed to proactively
                identify and address vulnerabilities across your digital
                ecosystem. By combining manual testing, automated tools, and
                industry best practices, we help enterprises strengthen their
                security posture, ensure compliance, and enable secure digital
                transformation.
              </p>
              <p className="mb-4">
                From code to cloud, our end-to-end approach to penetration
                testing (VAPT), secure DevOps, and red teaming helps you detect
                threats early, respond faster, and build resilience at scale.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/SecurityTesting.png"
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
