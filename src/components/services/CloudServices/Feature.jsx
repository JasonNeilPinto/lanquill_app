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
                  description="Tailored for Every Stage of Cloud Transformation."
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
                    <h3 className="h5">Rehost (&quot;Lift & Shift&quot;)</h3>
                    <p className="mb-0">
                      Quickly migrate applications to the cloud with minimal
                      changes to infrastructure and code.
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
                    <h3 className="h5">Replatform</h3>
                    <p className="mb-0">
                      Make slight optimizations to your workloads to take
                      advantage of cloud-native services without rewriting the
                      core application.
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
                    <h3 className="h5">Repurchase</h3>
                    <p className="mb-0">
                      Transition from legacy solutions to new SaaS platforms
                      that meet your evolving business needs.
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
                    <h3 className="h5">Refactor (Re-architect)</h3>
                    <p className="mb-0">
                      Redesign applications to fully leverage cloud-native
                      capabilities, microservices, and containerization for
                      improved performance and agility.
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
                    <h3 className="h5">Retire</h3>
                    <p className="mb-0">
                      Identify and eliminate obsolete or redundant applications
                      to simplify your cloud environment and reduce costs.
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
                    <h3 className="h5">Retain</h3>
                    <p className="mb-0">
                      Keep certain applications on-premises due to regulatory,
                      technical, or strategic reasons—integrating them
                      seamlessly into your hybrid cloud strategy.
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
