import React from "react";

const ContactBox = () => {
  return (
    <>
      <section className="contact-promo ptb-120">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-lg-4 col-md-6 mt-4 mt-lg-0">
              <div className="contact-us-promo p-5 bg-white rounded-custom custom-shadow text-center d-flex flex-column h-100">
                <span className="fad fa-comment-alt-lines fa-3x text-primary"></span>
                <div className="contact-promo-info mb-4">
                  <h5>Chat with us</h5>
                  <p>
                    We&apos;ve got live Social Experts waiting to help you{" "}
                    <strong>monday to friday</strong> from{" "}
                    <strong>9am to 5pm IST.</strong>
                  </p>
                </div>
                <a
                  href="mailto:hellothemetags@gmail.com"
                  className="btn btn-link mt-auto"
                >
                  Chat with us
                </a>
              </div>
            </div>
            <div className="col-lg-4 col-md-6 mt-4 mt-lg-0">
              <div className="contact-us-promo p-5 bg-white rounded-custom custom-shadow text-center d-flex flex-column h-100">
                <span className="fad fa-envelope fa-3x text-primary"></span>
                <div className="contact-promo-info mb-4">
                  <h5>Email Us</h5>
                  <p>
                    Simple drop us an email at{" "}
                    <strong>support@netanalytiks.com</strong>
                    and you&apos;ll receive a reply within 24 hours
                  </p>
                </div>
                <a
                  href="mailto:support@netanalytiks.com"
                  className="btn btn-primary mt-auto"
                >
                  Email Us
                </a>
              </div>
            </div>
            <div className="col-lg-4 col-md-6 mt-4 mt-lg-0">
              <div className="contact-us-promo p-5 bg-white rounded-custom custom-shadow text-center d-flex flex-column h-100">
                <span className="fad fa-phone fa-3x text-primary"></span>
                <div className="contact-promo-info mb-4">
                  <h5>Give us a call</h5>
                  <p>
                    Give us a ring.Our Experts are standing by{" "}
                    <strong>monday to friday</strong> from{" "}
                    <strong>9am to 5pm IST.</strong>
                  </p>
                </div>
                <a href="tel:+91 73383 36729" className="btn btn-link mt-auto">
                  +91 73383 36729
                </a>
              </div>
            </div>
          </div>

          {/* Office Address & Map */}
          <div className="row justify-content-center mt-5">
            <div className="col-lg-6 col-md-12 d-flex flex-column gap-4">
              <div className="contact-us-promo p-3 bg-white rounded-custom custom-shadow text-center d-flex flex-column">
                <span className="fad fa-location-dot fa-3x text-primary"></span>
                <div className="contact-promo-info">
                  <h5>Corporate office</h5>
                  <p>
                    <strong>NetAnalytiks Technologies Limited</strong> <br />
                    91springboard,
                    <br />
                    Gopala Krishna Complex 45/3,
                    <br />
                    Residency Road, Mahatma Gandhi Rd,
                    <br />
                    Bengaluru, Karnataka, India – 560025
                  </p>
                </div>
              </div>

              <div className="contact-us-promo p-3 bg-white rounded-custom custom-shadow text-center d-flex flex-column">
                <span className="fad fa-location-dot fa-3x text-primary"></span>
                <div className="contact-promo-info">
                  <h5>USA</h5>
                  <p>
                    375 Rantoul Street 209,
                    <br />
                    Beverly, Massachusetts - 01915, USA
                  </p>
                </div>
              </div>
            </div>

            <div className="col-lg-6 col-md-12">
              <div
                className="h-100 contact-us-promo p-0 bg-white rounded-custom custom-shadow"
                style={{ height: "100%" }}
              >
                <iframe
                  src="https://maps.google.com/maps?q=91springboard,%204th%20Floor,%20Gopala%20Krishna%20Complex%2045/3,%20Residency%20Road,%20Mahatma%20Gandhi%20Rd,%20Bengaluru,%20Karnataka%20560025&t=&z=9&ie=UTF8&iwloc=&output=embed"
                  width="100%"
                  height="100%"
                  style={{
                    border: 0,
                    display: "block",
                    minHeight: "100%",
                  }}
                  allowFullScreen=""
                  loading="lazy"
                  referrerPolicy="no-referrer-when-downgrade"
                  title="Google Maps Location"
                ></iframe>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default ContactBox;
