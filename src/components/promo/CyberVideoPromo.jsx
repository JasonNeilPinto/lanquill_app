import React from "react";

const CyberVideoPromo = () => {
  return (
    <section className="counter-with-video pt-60 pb-120">
      <div className="container">
        <div className="row">
          <div className="col-lg-6">
            <div
              className="cyber-video-img"
              style={{
                background:
                  "url('/img/promo-bg.png')no-repeat center top/cover",
              }}
            ></div>
          </div>
          <div className="col-lg-6">
            <div className="mt-5 mt-lg-0">
              <div className="section-heading mb-5">
                <h5 className="h6 text-primary">Partner with Us</h5>
                <h2>
                  Your Trusted IT, Gen AI, Cybersecurity and IT Consulting
                  Provider
                </h2>
                <p>
                  We deliver end-to-end solutions that help you innovate faster,
                  make smarter decisions, and scale with confidence.
                </p>
              </div>
              <div className="row">
                <div className="col-lg-6 col-md-6">
                  <div className="bg-white p-4 cyber-count-box mb-30 mb-lg-0">
                    <h2 className="text-primary">38+</h2>
                    <h5 className="h-6">Happy Clients</h5>
                    <p>
                      Driving measurable business impact through future-ready
                      IT, analytics, and security solutions.
                    </p>
                  </div>
                </div>
                <div className="col-lg-6 col-md-6">
                  <div className="bg-white p-4 cyber-count-box">
                    <h2 className="text-primary">170+</h2>
                    <h5 className="h-6">Success Projects</h5>
                    <p>
                      Delivering tailored strategies and transformative outcomes
                      for diverse industries worldwide.
                    </p>
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

export default CyberVideoPromo;
