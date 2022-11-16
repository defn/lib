package main

input: {
	terraform: {
		organization: "defn"
		workspace:    "workspaces"
	}

	organizations: [N=string]: name: N
	organizations: {
		_aws_admins: [{
			name:  "defn"
			email: "iam@defn.sh"
		}]

		_full_accounts: ["net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"]
		_env_accounts: ["net", "lib", "hub"]
		_ops_accounts: ["ops"]

		defn: {
			region: "us-west-2",
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: _ops_accounts
			admins:   _aws_admins
		}
		gyre: {
			region:   "us-east-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: _ops_accounts
			admins:   _aws_admins
		}
		curl: {
			region:   "us-west-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: _env_accounts
			admins:   _aws_admins
		}
		coil: {
			region:   "us-east-1"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: _env_accounts
			admins:   _aws_admins
		}
		helix: {
			region:   "us-east-2"
			prefix:   "aws-"
			domain:   "defn.sh"
			accounts: _full_accounts
			admins:   _aws_admins
		}
		spiral: {
			region:   "us-west-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: _full_accounts
			admins:   _aws_admins
		}
	}
}
