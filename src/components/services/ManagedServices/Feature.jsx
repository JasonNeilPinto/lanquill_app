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
                  description="Reliable IT Operations with End-to-End Managed Services."
                  centerAlign
                />
              )}
            </div>
          </div>
          <div className="row">
            <div className="col-12">
              <div className="feature-grid-ms">
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
                    <h3 className="h5">Managed Global Service Desk</h3>
                    <p className="mb-0">
                      24x7 multilingual support that ensures rapid incident
                      resolution, user support, and efficient ticket management.
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
                    <h3 className="h5">Managed Application Support Services</h3>
                    <p className="mb-0">
                      Ongoing monitoring, maintenance, and performance
                      optimization for enterprise applications to ensure
                      business continuity.
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
                    <h3 className="h5">
                      Managed Infrastructure Support Services
                    </h3>
                    <p className="mb-0">
                      Complete lifecycle management of servers, networks, and
                      endpoints—boosting uptime and operational resilience.
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
                    <h3 className="h5">Managed Security Services</h3>
                    <p className="mb-0">
                      End-to-end security management including threat detection,
                      vulnerability management, and compliance support to keep
                      your digital assets protected.
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
