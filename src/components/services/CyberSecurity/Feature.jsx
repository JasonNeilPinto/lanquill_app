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
                  description="Secure your business across every layer of digital interaction. From proactive monitoring to AI-specific safeguards, our cybersecurity portfolio addresses critical risk points to fortify your digital presence."
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
                    <i className="far fa-cloud icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Cloud & Infrastructure Security </h3>
                    <p className="mb-0">
                      Cloud Security Posture Management, Zero Trust Network
                      Access, SASE, and IDPS solutions to protect hybrid
                      infrastructures and enforce secure, policy-driven access.
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
                    <i className="fas fa-microchip icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Hardware & Data Security</h3>
                    <p className="mb-0">
                      Robust protection through Hardware Security Modules,
                      cryptographic key management, and advanced data encryption
                      and loss prevention strategies.
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
                    <i className="far fa-brain icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">AI & LLM Security </h3>
                    <p className="mb-0">
                      Cutting-edge protection against AI-specific
                      vulnerabilities with adversarial ML defense, prompt
                      injection protection, AI bias management, and secure
                      development protocols.
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
                    <i className="far fa-fingerprint icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Identity & Access Management</h3>
                    <p className="mb-0">
                      Seamless and secure access with Federated IAM, Privileged
                      Access controls, AI-powered adaptive authentication, and
                      multi-factor enforcement.
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
                    <i className="fal fa-exclamation-triangle icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Risk Management & Compliance</h3>
                    <p className="mb-0">
                      Comprehensive GRC solutions, regulatory compliance,
                      DevSecOps practices, and offensive testing through Red
                      Teaming & Penetration Testing services.
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
                    <i className="fal fa-shield-check icon-sm text-white"></i>
                  </div>
                  <div className="feature-content">
                    <h3 className="h5">Security Analytics</h3>
                    <p className="mb-0">
                      Integrated threat visibility via SIEM, SOAR, threat
                      intelligence, behavior analytics, and AI-powered Security
                      Operations Center for real-time response.
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
