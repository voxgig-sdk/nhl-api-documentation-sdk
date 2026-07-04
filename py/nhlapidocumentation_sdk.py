# NhlApiDocumentation SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import NhlApiDocumentationUtility
from core.spec import NhlApiDocumentationSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import NhlApiDocumentationBaseFeature
from features import _make_feature


class NhlApiDocumentationSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = NhlApiDocumentationUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features from config.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            feature_items = vs.items(feature_opts)
            if feature_items is not None:
                for item in feature_items:
                    fname = item[0]
                    fopts = helpers.to_map(item[1])
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return NhlApiDocumentationUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = NhlApiDocumentationSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    @property
    def conference(self):
        """Idiomatic facade: client.conference.list() / client.conference.load({"id": ...})."""
        from entity.conference_entity import ConferenceEntity
        cached = getattr(self, "_conference", None)
        if cached is None:
            cached = ConferenceEntity(self, None)
            self._conference = cached
        return cached

    def Conference(self, data=None):
        # Deprecated: use client.conference instead.
        from entity.conference_entity import ConferenceEntity
        return ConferenceEntity(self, data)


    @property
    def division(self):
        """Idiomatic facade: client.division.list() / client.division.load({"id": ...})."""
        from entity.division_entity import DivisionEntity
        cached = getattr(self, "_division", None)
        if cached is None:
            cached = DivisionEntity(self, None)
            self._division = cached
        return cached

    def Division(self, data=None):
        # Deprecated: use client.division instead.
        from entity.division_entity import DivisionEntity
        return DivisionEntity(self, data)


    @property
    def game(self):
        """Idiomatic facade: client.game.list() / client.game.load({"id": ...})."""
        from entity.game_entity import GameEntity
        cached = getattr(self, "_game", None)
        if cached is None:
            cached = GameEntity(self, None)
            self._game = cached
        return cached

    def Game(self, data=None):
        # Deprecated: use client.game instead.
        from entity.game_entity import GameEntity
        return GameEntity(self, data)


    @property
    def player(self):
        """Idiomatic facade: client.player.list() / client.player.load({"id": ...})."""
        from entity.player_entity import PlayerEntity
        cached = getattr(self, "_player", None)
        if cached is None:
            cached = PlayerEntity(self, None)
            self._player = cached
        return cached

    def Player(self, data=None):
        # Deprecated: use client.player instead.
        from entity.player_entity import PlayerEntity
        return PlayerEntity(self, data)


    @property
    def player_stat(self):
        """Idiomatic facade: client.player_stat.list() / client.player_stat.load({"id": ...})."""
        from entity.player_stat_entity import PlayerStatEntity
        cached = getattr(self, "_player_stat", None)
        if cached is None:
            cached = PlayerStatEntity(self, None)
            self._player_stat = cached
        return cached

    def PlayerStat(self, data=None):
        # Deprecated: use client.player_stat instead.
        from entity.player_stat_entity import PlayerStatEntity
        return PlayerStatEntity(self, data)


    @property
    def roster(self):
        """Idiomatic facade: client.roster.list() / client.roster.load({"id": ...})."""
        from entity.roster_entity import RosterEntity
        cached = getattr(self, "_roster", None)
        if cached is None:
            cached = RosterEntity(self, None)
            self._roster = cached
        return cached

    def Roster(self, data=None):
        # Deprecated: use client.roster instead.
        from entity.roster_entity import RosterEntity
        return RosterEntity(self, data)


    @property
    def schedule(self):
        """Idiomatic facade: client.schedule.list() / client.schedule.load({"id": ...})."""
        from entity.schedule_entity import ScheduleEntity
        cached = getattr(self, "_schedule", None)
        if cached is None:
            cached = ScheduleEntity(self, None)
            self._schedule = cached
        return cached

    def Schedule(self, data=None):
        # Deprecated: use client.schedule instead.
        from entity.schedule_entity import ScheduleEntity
        return ScheduleEntity(self, data)


    @property
    def standing(self):
        """Idiomatic facade: client.standing.list() / client.standing.load({"id": ...})."""
        from entity.standing_entity import StandingEntity
        cached = getattr(self, "_standing", None)
        if cached is None:
            cached = StandingEntity(self, None)
            self._standing = cached
        return cached

    def Standing(self, data=None):
        # Deprecated: use client.standing instead.
        from entity.standing_entity import StandingEntity
        return StandingEntity(self, data)


    @property
    def team(self):
        """Idiomatic facade: client.team.list() / client.team.load({"id": ...})."""
        from entity.team_entity import TeamEntity
        cached = getattr(self, "_team", None)
        if cached is None:
            cached = TeamEntity(self, None)
            self._team = cached
        return cached

    def Team(self, data=None):
        # Deprecated: use client.team instead.
        from entity.team_entity import TeamEntity
        return TeamEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None):
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk
