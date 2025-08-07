import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">Cloud Services</h2>
              <p className="mb-4">
                Our Cloud Services empower enterprises to modernize
                infrastructure, reduce operational complexity, and accelerate
                time-to-value through a strategic and customized cloud journey.
              </p>
              <p className="mb-4">
                Whether you&apos;re migrating legacy applications, optimizing
                workloads, or adopting multi-cloud and hybrid architectures, we
                support you at every step. Our cloud experts ensure secure,
                scalable, and cost-effective solutions tailored to your unique
                business goals.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/CloudServices.png"
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
