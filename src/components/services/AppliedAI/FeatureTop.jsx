import React from "react";
// import { Link } from "react-router-dom";

export default function FeatureTop() {
  return (
    <section className="payment-feature-img pt-120">
      <div className="container">
        <div className="row align-items-center">
          <div className="col-lg-5">
            <div className="payment-feature-img-left mb-5 mb-lg-0">
              <h2 className="mb-4">Applied AI</h2>
              <p className="mb-4">
                Our enterprise-grade AI solutions are designed to turn raw data
                into smart actions—optimizing processes, enhancing
                decision-making, and elevating user experiences. With expertise
                spanning machine learning (ML), natural language processing
                (NLP), and computer vision, our Applied AI services enable
                intelligent systems that not only understand but act with
                precision and purpose.
              </p>
            </div>
          </div>
          <div className="col-lg-7">
            <div className="text-center position-relative mt-5 mt-lg-0">
              <div className="payment-feature-mockup position-relative p-2">
                <img
                  src="/img/services/feature-top/AppliedAI.png"
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
