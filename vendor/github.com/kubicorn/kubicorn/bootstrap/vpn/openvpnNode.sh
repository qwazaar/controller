# ------------------------------------------------------------------------------------------------------------------------
# We are explicitly not using a templating language to inject the values as to encourage the user to limit their
# use of templating logic in these files. By design all injected values should be able to be set at runtime,
# and the shell script real work. If you need conditional logic, write it in bash or make another shell script.
# ------------------------------------------------------------------------------------------------------------------------

apt-get update -y
apt-get install -y jq

VPN=$(cat /etc/kubicorn/cluster.json | jq -r '.components.vpn')

if [ -z $VPN ]; then
    sleep 7
    export VPN=$(cat /etc/kubicorn/cluster.json | jq -r '.components.vpn')
fi

if $VPN; then
    PRIVATE_IP=$(curl http://169.254.169.254/metadata/v1/interfaces/private/0/ipv4/address)

    # OpenVPN
    apt-get install -y openvpn

    OPENVPN_CONF=$(cat /etc/kubicorn/cluster.json | jq -r '.values.itemMap.INJECTEDCONF')

    echo -e ${OPENVPN_CONF} > /etc/openvpn/clients.conf

    systemctl start openvpn@clients
    systemctl enable openvpn@clients
fi
