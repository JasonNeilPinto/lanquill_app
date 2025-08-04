import React from "react";
import HeroTitle from "../components/common/HeroTitle.jsx";

const DesktopHome = () => {
  return (
    <>
      <section
        className="hero-section ptb-120"
        style={{
          background:
            "url('/img/shape/dot-dot-wave-shape.svg')no-repeat bottom center",
        }}
      >
        <div className="container">
          <div className="row align-items-center justify-content-lg-between">
            <div className="col-xl-5 col-lg-5">
              <div
                className="hero-content-wrap text-center text-xl-start text-lg-start"
                data-aos="fade-right"
              >
                <HeroTitle
                  title="Powerful Solutions for Your Business"
                  desc="   Proactively coordinate quality quality vectors vis-a-vis
                  supply chains client-centric web services."
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-6 mt-4 mt-xl-0">
              <div
                className="hero-img-wrap position-relative"
                data-aos="fade-left"
              >
                <ul className="position-absolute animate-element parallax-element shape-service hide-medium">
                  <li className="layer" data-depth="0.03">
                    <img
                      src="/img/color-shape/image-1.svg"
                      alt="shape"
                      className="img-fluid position-absolute color-shape-1"
                    />
                  </li>
                  <li className="layer" data-depth="0.02">
                    <img
                      src="/img/color-shape/feature-2.svg"
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

                <div className="hero-img-wrap position-relative">
                  <div className="hero-screen-wrap">
                    <div className="phone-screen">
                      <img
                        src="/img/screen/phone-screen.png"
                        alt="hero"
                        className="position-relative img-fluid"
                      />
                    </div>
                    <div className="mac-screen">
                      <img
                        src="/img/screen/mac-screen.png"
                        alt="hero "
                        className="position-relative img-fluid rounded-custom"
                      />
                    </div>
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

export default DesktopHome;
