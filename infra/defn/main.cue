package main

aws_admins: [{
	name:  "defn"
	email: "iam@defn.sh"
}, {
	name:  "dgwyn"
	email: "david@defn.sh"
}]

full_accounts: ["ops", "net", "lib", "hub", "log", "sec", "pub", "dev", "dmz"]
env_accounts: ["ops", "net", "lib", "hub"]
ops_accounts: ["ops"]

input: {
	terraform: {
		organization: "defn"
		workspace:    "workspaces"
	}

	organizations: [N=string]: name: N
	organizations: {
		gyre: {
			region:   "us-east-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: ops_accounts
			admins:   aws_admins
		}
		curl: {
			region:   "us-west-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: env_accounts
			admins:   aws_admins
		}
		coil: {
			region:   "us-east-1"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: env_accounts
			admins:   aws_admins
		}
		helix: {
			region:   "us-east-2"
			prefix:   "aws-"
			domain:   "defn.sh"
			accounts: full_accounts
			admins:   aws_admins
		}
		spiral: {
			region:   "us-west-2"
			prefix:   "aws-"
			domain:   "defn.us"
			accounts: full_accounts
			admins:   aws_admins
		}
	}
}
