package eu.erdin.throttler

import io.gatling.core.Predef._
import io.gatling.http.Predef._

import scala.concurrent.duration._


class StressTest extends Simulation {

  val httpConf = http
    .baseURL("http://localhost:8080")

  setUp(
    scenario("most-important")
      .exec(http("proxy").get("/"))
      .pause(5,12)
      .exec(http("proxy").get("/"))
      .pause(5,12)
      .exec(http("proxy").get("/"))
      .pause(5,12)
      .exec(http("proxy").get("/"))
      .inject(rampUsersPerSec(10) to(200) during(2 minutes)).protocols(httpConf)
  )
}