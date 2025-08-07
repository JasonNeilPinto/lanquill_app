/* eslint-disable react/prop-types */
import React from "react";
import SectionTitle from "../../common/SectionTitle";

const Feature = ({ cardDark }) => {
  return (
    <>
      <section
        className={`feature-section ptb-120 ${
          cardDark ? "bg-dark" : "bg-light"
        }`}
      >
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-lg-10 col-md-10">
              {cardDark ? (
                <SectionTitle
                  subtitle="Services"
                  title="Best Services Grow Your Business Value"
                  description="Globally actualize cost effective with resource maximizing
                  leadership skills."
                  centerAlign
                  dark
                />
              ) : (
                <SectionTitle
                  subtitle="Services"
                  title="Explore our Services"
                  description="Ensure robust protection with our comprehensive security testing services"
                  centerAlign
                />
              )}
            </div>
          </div>
          <div className="row">
            <div className="col-12">
              <div className="feature-grid">
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="far fa-database icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Web Application VAPT</h3>
                    <p className="mb-0">
                      Simulate real-world attacks to uncover vulnerabilities in
                      web apps, APIs, and business logic using OWASP Top 10 and
                      beyond.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-lightbulb icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Mobile App VAPT</h3>
                    <p className="mb-0">
                      Assess mobile apps for platform-specific threats, data
                      leakage, and authentication flaws across iOS and Android
                      ecosystems.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-chart-line icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Infrastructure VAPT</h3>
                    <p className="mb-0">
                      Perform deep-dive testing of network, servers, cloud
                      infrastructure, and endpoints to reveal internal and
                      external security risks.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-file-chart-line icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Security Code Reviews</h3>
                    <p className="mb-0">
                      Conduct thorough static code analysis to identify insecure
                      coding practices and misconfigurations at the development
                      stage.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-exchange-alt icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">DevSecOps Integration</h3>
                    <p className="mb-0">
                      Shift security left with integrated DevSecOps pipelines,
                      enabling continuous security through automation and early
                      detection.
                    </p>
                  </div>
                </div>
                <div
                  className={`feature-card shadow-sm rounded-custom p-5 ${
                    cardDark
                      ? "bg-custom-light promo-border-hover border border-2 border-light text-white"
                      : "bg-white"
                  }`}
                  data-aos="fade-up"
                  data-aos-delay="50"
                >
                  <div
                    className="
                      icon-box
                      d-inline-block
                      rounded-circle
                      bg-primary-soft
                      mb-32
                    "
                  >
                    <i className="fal fa-compass icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Red Teaming</h3>
                    <p className="mb-0">
                      Simulate advanced persistent threats (APT) and adversarial
                      behavior to test real-time response, detection, and
                      resilience.
                    </p>
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

export default Feature;
