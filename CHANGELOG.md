# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Next version

* Automatically redraw the dashboard when size of terminal change.
* Fix errors linked to Google JWT files.
* Command list - list the config available

## [0.4.2] - 2020-02-25

### Updated

* Possibility to overwrite the URL from the monitoring service in a monitoring widget (allows to monitor one address by box widget).
* Fix a small display bug (forgot to delete a useless debug output... oops).

## [0.4.1] - 2020-11-24

### Updated

* Fix filepaths when using DevDash on Windows

## [0.4.0] - 2020-06-08

### Added

* Google Analytics widgets:
    * ga.bar_countries
    * ga.bar_devices

* Remote host / local host service (grab data from command line via SSH on remote or local host)
    * rh.box_uptime 
    * rh.box_load
    * rh.box_net_io
    * rh.box_disk_io
    * rh.bar_memory
    * rh.gauge_cpu_rate
    * rh.gauge_memory_rate
    * rh.gauge_swap_rate
    * rh.bar_rates
    * rh.table_disk
    * rh.box
    * rh.gauge
    * rh.table
    * rh.bar

* Possibility to load dashboard configurations from `$XDG_CONFIG_HOME` only using filename.
    * Example: `devdash -c dashboard` if there is a file `$XDG_CONFIG_HOME/devdash/dashboard.yml/json`

### Updated 

* Increase performances

### Breaking Changes

* `devdash -config dashboad-config.yml` is not valid anymore. Replaced by `devdash -c dashboard-config.yml` or `devdash --config dashboard-config.yml`

## [0.3.0] - 2020-01-14

### Added

* Add Git service
    * Add widget table branches - display information about git branches
* Add Display service
    * Add widget box - display a widget box containing some text
* Add Travis CI service
    * Add Travis builds widget - display your last Travis CI builds
* Add Feedly service
    * Add Feedly subscribers widget - display number of Feedly subscribers for your website

## [0.2.0] - 2019-10-05

### Added

* Add Github widgets
  * Display count stars overtime
  * Display count commits overtime
  * Display issues
  * Display repositories in table with information
  * Display last week traffic on Github page

* Add `color` options to have same color for border, title and everything color related for one widget
* Add themes to simplify the configuration - possibility to use same options defined once, for multiple widgets

* Add possibility to hot reload any dashboard via a keystroke - no need to restart DevDash when changing a dashboard configuration

* Create the [official DevDash website](https://thedevdash.com)

### Updated 

* Replace `title_options` by `name_options` for project's config (breaking change)

## [0.1.1] - 2019-07-21

### Added

* use goreleaser for relases

## [0.1.0] - 2019-05-28

### Added

* Write README documentation
* Add Github widgets
  * github.box_stars
  * github.box_watchers
  * github.box_open_issues
  * github.table_branches
  * github.table_issues
  * github.table_repositories
* Add Github API
* Google Search Console widgets
  * gsc.table_pages
  * gsc.table_queries
  * gsc.table
* Add Google Search Console API
* Create ToTime library
* Google Analytics widgets:
  * ga.box_real_time
  * ga.box_total
  * ga.bar_sessions
  * ga.bar_bounces
  * ga.bar_users
  * ga.bar_returning
  * ga.bar_pages
  * ga.bar
  * ga.bar_new_returning
  * ga.table_pages
  * ga.table_traffic_sources
  * ga.table
* Google Analytics service
* Google Analytics API
* Monitoring service
* Dashboard refreshing system 
* Projects / Services / Widget system
* Display and Grid system
* YAML configuration system

[0.1.1]: https://github.com/Phantas0s/devdash/releases/tag/v0.1.1
[0.1.0]: https://github.com/Phantas0s/devdash/releases/tag/v0.1.0
