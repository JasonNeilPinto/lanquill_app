import React from "react";
import SectionTitle from "../common/SectionTitle";

const OurStory = () => {
  return (
    <>
      <section
        className="our-story-section pt-60 pb-120"
        style={{
          background:
            "url('/img/shape/dot-dot-wave-shape.svg')no-repeat left bottom",
        }}
      >
        <div className="container">
          <div className="row justify-content-between">
            <div className="col-lg-5 col-md-12 order-lg-1">
              <div className="section-heading sticky-sidebar">
                <SectionTitle
                  subtitle="Our Story"
                  title="Who We Are"
                  description="NetAnalytiks is a multinational IT consulting and delivery services organization, crafting innovative solutions with new-age technologies. Our flexible engagement models—consulting, fixed-scope delivery, and SLA-based managed services—help businesses adapt and thrive."
                />
                <SectionTitle
                  title="What We Do"
                  description="We bring agility, precision, and deep expertise to every project—driving digital transformation, Generative AI adoption, cybersecurity, and analytics success. With proven products, skilled teams, and a commitment to rapid value creation, we deliver high-quality solutions that ensure measurable impact and long-term growth."
                />
              </div>
            </div>
            <div className="col-lg-6 col-md-12 order-lg-0">
              <div className="story-grid-wrapper position-relative">
                {/* <!--animated shape start--> */}
                <ul className="position-absolute animate-element parallax-element shape-service z--1">
                  <li className="layer" data-depth="0.02">
                    <img
                      src="/img/color-shape/image-2.svg"
                      alt="shape"
                      className="img-fluid position-absolute color-shape-2 z-5"
                    />
                  </li>
                  <li className="layer" data-depth="0.03">
                    <img
                      src="/img/color-shape/feature-3.svg"
                      alt="shape"
                      className="img-fluid position-absolute color-shape-3"
                    />
                  </li>
                </ul>
                {/* <!--animated shape end--> */}
                <div className="story-grid rounded-custom bg-dark overflow-hidden position-relative">
                  <div className="story-item bg-white border">
                    <h3 className="display-5 fw-bold mb-1 text-primary">
                      120+
                    </h3>
                    <h6 className="mb-0">Skilled Experts</h6>
                  </div>
                  <div className="story-item bg-light border">
                    <h3 className="display-5 fw-bold mb-1 text-warning">
                      10+ Years
                    </h3>
                    <h6 className="mb-0">In Business</h6>
                  </div>
                  <div className="story-item bg-light border">
                    <h3 className="display-5 fw-bold mb-1 text-danger">38+</h3>
                    <h6 className="mb-0">Clients Worldwide</h6>
                  </div>
                  <div className="story-item bg-white border">
                    <h3 className="display-5 fw-bold mb-1 text-primary">
                      170+
                    </h3>
                    <h6 className="mb-0">Projects Completed</h6>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default OurStory;
